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

    leveldbDir := os.Getenv("LEVELDB_DIR")

    if leveldbDir == "" {
        leveldbDir = "leveldb-data"
    }

    dbPath := filepath.Join(leveldbDir)

    fmt.Println("Connecting to LevelDB at", dbPath)
    db, err := leveldb.OpenFile(dbPath, nil)
    if err != nil {
        log.Fatalf("Failed to connect to LevelDB: %v", err)
    }
    defer db.Close()
    fmt.Println("Connected to LevelDB")

    fmt.Println("Hello, World!")
}
