package models

import (
	"fmt"

	customErrors "github.com/kindaqt/assignment2/errors"
)

// Cache holds data in memory
type Cache struct {
	Values map[string][]byte
}

// NewCachePersistence returns a Persistence interface
func NewCachePersistence() Persistence {
	return NewCache()
}

// NewCache returns an instance of CacheInterface which has all the Persistence interface functionality plus additional functions by way of interface composition
func NewCache() CacheInterface {
	values := make(map[string][]byte)
	return &Cache{
		Values: values,
	}
}

// Put updates or replaces resources in the repository based on the existence of said resource
func (p *Cache) Put(key string, value []byte) error {
	// Store Values
	p.Values[key] = value

	return nil
}

// Get retrieves a resource based on the key
func (p *Cache) Get(key string) ([]byte, error) {
	// Get Values
	b, ok := p.Values[key]
	if !ok {
		return nil, customErrors.TemporaryError{fmt.Sprintf("Error while getting %v", key)}
	}

	return b, nil
}

// Flush deletes a record from cache
func (p *Cache) Flush(key string) {
	delete(p.Values, key)
}
