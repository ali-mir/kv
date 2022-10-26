package storage_engine

import "testing"


func TestInsertAndLookup(t *testing.T) {
	s := NewStorageEngine()
	s.Insert("hello", "world")
	value := s.Lookup("hello")
	if value != "world" {
		panic("Expected {\"hello\": \"world\"}")
	}
}

func TestDelete(t *testing.T) {
	s := NewStorageEngine()
	s.Insert("hello", "world")
	s.Delete("hello")
	value := s.Lookup("hello")
	if value != "" {
		panic("Expected nil for key \"hello\"")
	}
}