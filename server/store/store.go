package store

import (
	"log"
	"sync"
)

var m map[string]string
var mutex sync.Mutex

func Initialize() {
	m = make(map[string]string)
}

func Insert(key, value string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	log.Printf("Inserting key {%s} with value {%s}", key, value)
	m[key] = value
	return true
}

func Lookup(key string) string {
	log.Printf("Looking up key {%s}", key)
	return m[key]
}

func Delete(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	log.Printf("Deleting key {%s}", key)
	delete(m, key)
	return true
}
