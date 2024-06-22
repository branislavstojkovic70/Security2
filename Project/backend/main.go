package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
	err := godotenv.Load()
    leveldbHost := os.Getenv("LEVELDB_HOST")
    leveldbPort := os.Getenv("LEVELDB_PORT")
	fmt.Println(leveldbHost)
	fmt.Println(leveldbPort)

    if leveldbHost == "" {
        leveldbHost = "leveldb-server"
    }
    if leveldbPort == "" {
        leveldbPort = "2012"
    }

    dbPath := fmt.Sprintf("http://%s:%s", leveldbHost, leveldbPort)
    fmt.Println("Connecting to LevelDB at", dbPath)
    db, err := leveldb.OpenFile(dbPath, nil)
    if err != nil {
        log.Fatalf("Failed to connect to LevelDB: %v", err)
    }
    defer db.Close()
    fmt.Println("Connected to LevelDB")

    fmt.Println("Hello, World!")
}
