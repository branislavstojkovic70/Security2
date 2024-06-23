package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/hashicorp/consul/api"
    "strings"
)

type ACL struct {
    Object   string `json:"object"`
    Relation string `json:"relation"`
    User     string `json:"user"`
}

type Namespace struct {
    Namespace string                `json:"namespace"`
    Relations map[string]Relation `json:"relations"`
}

type Relation struct {
    Union []map[string]interface{} `json:"union"`
}

var (
    db *leveldb.DB
    consulClient *api.Client
)

func initDB(leveldbDir string) error {
    var err error
    db, err = leveldb.OpenFile(leveldbDir, nil)
    return err
}

func closeDB() {
    db.Close()
}

func initConsul() error {
    var err error
    consulConfig := api.DefaultConfig()
    consulClient, err = api.NewClient(consulConfig)
    return err
}

func handleACL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var acl ACL
    err := json.NewDecoder(r.Body).Decode(&acl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    key := fmt.Sprintf("%s#%s@%s", acl.Object, acl.Relation, acl.User)
    err = db.Put([]byte(key), []byte{}, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func checkACL(object, relation, user string) (bool, error) {
    // Prvo proverite direktnu relaciju
    key := fmt.Sprintf("%s#%s@%s", object, relation, user)
    _, err := db.Get([]byte(key), nil)
    if err == nil {
        return true, nil
    }

    // Ako relacija nije direktna, proverite computed_userset
    namespaceKey := "" + strings.Split(object, ":")[0]
    kv, _, err := consulClient.KV().Get(namespaceKey, nil)
    if err != nil || kv == nil {
        return false, fmt.Errorf("namespace not found")
    }

    var namespace Namespace
    err = json.Unmarshal(kv.Value, &namespace)
    if err != nil {
        return false, fmt.Errorf("failed to unmarshal namespace: %v", err)
    }

    relationDef, exists := namespace.Relations[relation]
    if !exists {
        return false, fmt.Errorf("relation not found")
    }

    for _, rel := range relationDef.Union {
        if _, exists := rel["this"]; exists {
            // Do nothing, already checked the direct relation
        } else if computedUserset, exists := rel["computed_userset"]; exists {
            computedRel := computedUserset.(map[string]interface{})
            computedRelation := computedRel["relation"].(string)
            if result, err := checkACL(object, computedRelation, user); result && err == nil {
                return true, nil
            }
        }
    }

    return false, nil
}

func handleACLCheck(w http.ResponseWriter, r *http.Request) {
    object := r.URL.Query().Get("object")
    relation := r.URL.Query().Get("relation")
    user := r.URL.Query().Get("user")

    authorized, err := checkACL(object, relation, user)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]bool{"authorized": authorized})
}

func handleNamespace(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var namespace Namespace
    err := json.NewDecoder(r.Body).Decode(&namespace)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    data, err := json.Marshal(namespace)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    kv := consulClient.KV()
    p := &api.KVPair{Key: namespace.Namespace, Value: data}
    _, err = kv.Put(p, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func handleGetNamespace(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    namespace := r.URL.Query().Get("namespace")
    key := "" + namespace
    kv, _, err := consulClient.KV().Get(key, nil)
    if err != nil || kv == nil {
        http.Error(w, "Namespace not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(kv.Value)
}
