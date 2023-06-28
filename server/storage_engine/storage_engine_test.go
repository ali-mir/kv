package storage_engine

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestInsertAndLookup(t *testing.T) {
	s := NewStorageEngine()
	s.Insert("hello", "world")
	value := s.Lookup("hello")
	if value != "world" {
		panic("Expected {\"hello\": \"world\"}")
	}
	assert.Equal(t, "world", value, "Expected {\"hello\": \"world\"}")

}

func TestDelete(t *testing.T) {
	s := NewStorageEngine()
	s.Insert("hello", "world")
	s.Delete("hello")
	value := s.Lookup("hello")
	assert.Equal(t, "", value, "Expected nil for key \"hello\"")
}