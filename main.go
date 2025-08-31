package main

import (
	"fmt"
    "encoding/gob"
    "os"
    "sync"
)

func main() {
	fmt.Println(MagentaBold, "The MyRedis Database!", Normal)
}

type Database struct {
    data map[string]any
    lock sync.RWMutex
}

func NewDatabase() *Database {
    return &Database{
        data: make(map[string]any)
    }
}

func (db *Database) Set(key string, value any) {
    db.lock.Lock()
    defer db.lock.Unlock()
    db.data[key] = value
}

func (db *Database) Get(key string) (any, bool) {
    db.lock.RLock()
    defer db.lock.RUnlock()
    value, ok := db.data[key]
    return value, ok
}
