package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "github.com/joho/godotenv"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    leveldbHost := os.Getenv("LEVELDB_HOST")
    leveldbPort := os.Getenv("LEVELDB_PORT")

    if leveldbHost == "" {
        leveldbHost = "leveldb-server"
    }
    if leveldbPort == "" {
        leveldbPort = "2012"
    }

    dbDir := "leveldb"
    dbPath := filepath.Join(dbDir, fmt.Sprintf("%s:%s", leveldbHost, leveldbPort))

    fmt.Println("Connecting to LevelDB at", dbPath)
    db, err := leveldb.OpenFile(dbPath, nil)
    if err != nil {
        log.Fatalf("Failed to connect to LevelDB: %v", err)
    }
    defer db.Close()
    fmt.Println("Connected to LevelDB")

    fmt.Println("Hello, World!")
}
