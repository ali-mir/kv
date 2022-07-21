package store

import (
	"log"
	"sync"
)

// thread-safe reads and writes
type Store struct {
	m map[string]string
	mutex sync.RWMutex
}

func NewStore() *Store {
	s := Store{}
	s.m = make(map[string]string)
	return &s
}

func (s *Store) Insert(key, value string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Inserting key {%s} with value {%s}", key, value)
	s.m[key] = value
	return true
}

func (s *Store) Lookup(key string) string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	log.Printf("Looking up key {%s}", key)
	return s.m[key]
}

func (s *Store) Delete(key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Deleting key {%s}", key)
	delete(s.m, key)
	return true
}
