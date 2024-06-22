package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "os"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
    leveldbHost := os.Getenv("LEVELDB_HOST")
    leveldbPort := os.Getenv("LEVELDB_PORT")
    consulHost := os.Getenv("CONSUL_HOST")
    consulPort := os.Getenv("CONSUL_PORT")

    if leveldbHost == "" {
        leveldbHost = "leveldb-server"
    }
    if leveldbPort == "" {
        leveldbPort = "2012"
    }
    if consulHost == "" {
        consulHost = "consul"
    }
    if consulPort == "" {
        consulPort = "8500"
    }

    // Testiranje konekcije na LevelDB
    dbPath := fmt.Sprintf("http://%s:%s", leveldbHost, leveldbPort)
    fmt.Println("Connecting to LevelDB at", dbPath)
    db, err := leveldb.OpenFile(dbPath, nil)
    if err != nil {
        log.Fatalf("Failed to connect to LevelDB: %v", err)
    }
    defer db.Close()
    fmt.Println("Connected to LevelDB")

    // Testiranje konekcije na Consul
    consulURL := fmt.Sprintf("http://%s:%s/v1/agent/self", consulHost, consulPort)
    fmt.Println("Connecting to Consul at", consulURL)
    resp, err := http.Get(consulURL)
    if err != nil {
        log.Fatalf("Failed to connect to Consul: %v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response from Consul: %v", err)
    }
    fmt.Println("Connected to Consul, response:", string(body))

    // Vaš glavni logički kod dolazi ovde
    fmt.Println("Hello, World!")
}
