package main

import (
	"encoding/gob"
	"fmt"
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
		data: make(map[string]any),
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

func (db *Database) Persis(filename string) error {
	db.lock.RLock()
	defer db.lock.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(db.data); err != nil {
		return err
	}

	return nil
}

func (db *Database) Load(filename string) error {
    db.lock.Lock()
    defer db.lock.Unlock()

    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := gob.NewDecoder(file)
    if err := decoder.Decode(&db.data); err != nil {
        return err
    }

    return nil
}


