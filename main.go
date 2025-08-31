package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Println(Prompt("The MyRedis Database!"))
	db := NewDatabase()
	db.Set("key1", "value1")
	db.Set("key2", "value2")

	err := db.Persist("Database.gob")
	if err != nil {
		fmt.Println("Failed to persist database:", err)
		return
	}

	err = db.Load("Database.gob")
	if err != nil {
		fmt.Println("Failed to load databse:", err)
		return
	}

	value1, ok1 := db.Get("key1")
	fmt.Println(Out(fmt.Sprint("Value1: ", value1, "Exists: ", ok1)))
	value2, ok2 := db.Get("key2")
	fmt.Println(Out(fmt.Sprint("Value2: ", value2, "Exists: ", ok2)))
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

func (db *Database) Persist(filename string) error {
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
