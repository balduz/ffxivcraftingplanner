package cache

import (
	"sync"
)

// Storage provides an interface to get and set cached elements.
type Storage interface {
	Get(id int) (interface{}, bool)
	Set(id int, data interface{})
}

type memoryStorage struct {
	elements map[int]interface{}
	mu       *sync.RWMutex
}

// NewStorage creates memory storage.
func NewStorage() Storage {
	return &memoryStorage{
		elements: make(map[int]interface{}),
		mu:       &sync.RWMutex{},
	}
}

// Get a cached element from memory.
func (s memoryStorage) Get(id int) (interface{}, bool) {
	s.mu.RLock()
	elem, ok := s.elements[id]
	s.mu.RUnlock()

	return elem, ok
}

// Set saves data with key id to memory cache.
func (s memoryStorage) Set(id int, data interface{}) {
	s.mu.Lock()
	s.elements[id] = data
	s.mu.Unlock()
}
