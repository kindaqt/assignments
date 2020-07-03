package models

import (
	"testing"

	"github.com/google/uuid"
	customErrors "github.com/kindaqt/assignment2/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//////////////////////////////
// Setup
/////////////////////////////

// Test Suite for Shared Resources
type PersistenceCacheTestSuite struct {
	suite.Suite
	cache *Cache
}

// Setup Test Suite
func TestPersistenceCacheTestSuite(t *testing.T) {
	suite.Run(t, new(PersistenceCacheTestSuite))
}

// Setup before each test
func (s *PersistenceCacheTestSuite) SetupTest() {
	values := make(map[string][]byte)
	s.cache = &Cache{values}
}

////////////////////////////
// Tests
///////////////////////////

func TestNewCachePersistence(t *testing.T) {
	t.Log("It should create a new cache instance.")
	dataStore := NewCachePersistence()

	assert.IsType(t, &Cache{}, dataStore, "It should return a cache.")
	_, ok := dataStore.(Persistence)
	assert.Equal(t, true, ok, "NewCachePersistence() should return an object that implements the Persistence interface")
}

func (s *PersistenceCacheTestSuite) TestPut() {
	key := uuid.New().String()
	value := []byte{0, 1, 2}
	s.T().Logf("Put() should add the key/value pair (%s : %v) to the cache values.", key, value)

	err := s.cache.Put(key, value)
	s.NoError(err, "Put() should not return an error.")
	s.Equal(s.cache.Values[key], value)
}

func (s *PersistenceCacheTestSuite) TestGet() {
	key := uuid.New().String()
	value := []byte{0, 1, 2}
	s.cache.Values[key] = value
	s.T().Logf("Get() should return %v of the the key (%v).", value, key)
	b, err := s.cache.Get(key)
	s.NoError(err, "Get should not return an error.")
	s.Equal(b, value)
}

func (s *PersistenceCacheTestSuite) TestGetError() {
	s.T().Logf("Get() return an error when the requested key does not exist.")
	_, err := s.cache.Get("badkey")
	s.Error(err, "Get should return an error.")
	s.IsType(customErrors.TemporaryError{}, err, "Get should return a temporary error.")
}

func (s *PersistenceCacheTestSuite) TestFlush() {
	key := uuid.New().String()
	value := []byte{0, 1, 2}
	s.cache.Values[key] = value
	s.T().Logf("Flush() should remove the requested key from the cache.")
	s.cache.Flush(key)
	s.Nil(s.cache.Values[key], "")
}
