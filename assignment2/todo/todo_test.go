package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	customErrors "github.com/kindaqt/assignment2/errors"
	"github.com/kindaqt/assignment2/models"
	"github.com/kindaqt/assignment2/test/mocks/mock_models"
	"github.com/stretchr/testify/suite"
)

//////////////////////////////
// Setup
/////////////////////////////

// Test Suite for Shared Resources
type TodoTestSuite struct {
	suite.Suite
	mockPersistence  *mock_models.MockPersistence
	mockCache        *mock_models.MockCacheInterface
	cache            models.CacheInterface
	cachePersistence models.Persistence
	todoDAO          TodoDAOPersister
}

var runIntegrationTests bool

// Setup Test Suite
func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}

// Setup before each test
func (s *TodoTestSuite) SetupTest() {
	// Setup Mocks
	persitenceCtrl := gomock.NewController(s.T())
	s.mockPersistence = mock_models.NewMockPersistence(persitenceCtrl)
	cacheCtrl := gomock.NewController(s.T())
	s.mockCache = mock_models.NewMockCacheInterface(cacheCtrl)

	// Setting CacheActive based on ENV variable.
	os.Setenv("CACHE_ACTIVE", "false")
	cacheActive, err := strconv.ParseBool(os.Getenv("CACHE_ACTIVE"))
	s.NoError(err)

	s.todoDAO = TodoDAOPersister{
		DataStore:   s.mockPersistence,
		Cache:       s.mockCache,
		CacheActive: cacheActive,
	}
}

// Test Variables
var testTodo Todo = Todo{
	ID:      uuid.New().String(),
	Title:   "Test Title",
	Message: "This is a test message",
}

var temporaryError = customErrors.TemporaryError{"some temporary error"}

////////////////////////////
// Tests
///////////////////////////

func (s *TodoTestSuite) TestSaveCacheInactive() {
	s.T().Log("Save() should save a record to the data store.")

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)
	s.mockCache.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)
	s.mockCache.EXPECT().Flush(testTodo.ID).Times(0)

	// Test Save()
	err = s.todoDAO.Save(testTodo)
	s.NoError(err)
}

func (s *TodoTestSuite) TestSaveCacheActive() {
	s.T().Log("Save() should save a record to the cache and the data store.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)
	s.mockCache.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)
	s.mockCache.EXPECT().Flush(testTodo.ID).Times(1)

	// Test Save()
	err = s.todoDAO.Save(testTodo)
	s.NoError(err)
}

func (s *TodoTestSuite) TestSaveErrorCacheActive() {
	s.T().Log("Save() should retrun an error.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Put(testTodo.ID, expectedByteArray).Return(errors.New("some error")).Times(1)
	s.mockCache.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)
	s.mockCache.EXPECT().Flush(testTodo.ID).Times(1)

	// Test Save()
	s.Error(s.todoDAO.Save(testTodo))
}

func (s *TodoTestSuite) TestSaveErrorCacheInactive() {
	s.T().Log("Save() should retrun an error.")

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Put(testTodo.ID, expectedByteArray).Return(errors.New("some error")).Times(1)

	// Test Save()
	s.Error(s.todoDAO.Save(testTodo))
}

func (s *TodoTestSuite) TestSaveCustomErrorCacheActive() {
	s.T().Log("Save() should retrun an error.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	expectedError := customErrors.TemporaryError{"some temporary error"}

	s.mockPersistence.EXPECT().
		Put(testTodo.ID, expectedByteArray).
		Return(expectedError).
		Times(3)
	s.mockCache.EXPECT().Put(testTodo.ID, expectedByteArray).Return(nil).Times(1)

	// Test Save()
	s.Error(s.todoDAO.Save(testTodo))
}

func (s *TodoTestSuite) TestGetByIDCacheInactive() {
	s.T().Log("GetByID() should returns a Todo based on the ID.")

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(expectedByteArray, nil).Times(1)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.NoError(err, "GetByID() should not return an error.")
	s.Equal(testTodo, actualTodo, "GetByID() should return testTodo.")
}

func (s *TodoTestSuite) TestGetByIDCacheActive() {
	s.T().Log("GetByID() should returns a Todo based on the ID.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(expectedByteArray, nil).Times(1)
	s.mockCache.EXPECT().Get(testTodo.ID).Return(expectedByteArray, nil).Times(1)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.NoError(err, "GetByID() should not return an error.")
	s.Equal(testTodo, actualTodo, "GetByID() should return testTodo.")
}

func (s *TodoTestSuite) TestGetByIDErrorCacheInactive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Mock Expectations: return nil, error
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, temporaryError).Times(3)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.Error(err, "GetByID() should return an error when Get() returns an error.")
	s.Equal(Todo{}, actualTodo, "GetByID() should return an empty Todo.")
}

func (s *TodoTestSuite) TestGetByIDTemporaryErrorCacheInactive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Mock Expectations: return nil, error
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, temporaryError).Return(expectedByteArray, nil)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.NoError(err, "GetByID() should not return an error.")
	s.Equal(testTodo, actualTodo, "GetByID() should return testTodo.")
}

func (s *TodoTestSuite) TestGetByIDErrorCacheActive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations: return nil, error
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, errors.New("some error")).Times(1)
	s.mockCache.EXPECT().Get(testTodo.ID).Return(nil, nil).Times(1)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.Error(err, "GetByID() should return an error when Get() returns an error.")
	s.Equal(Todo{}, actualTodo, "GetByID() should return an empty Todo.")
}

func (s *TodoTestSuite) TestGetByIDCustomErrorCacheInactive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Mock Expectations: return nil, error
	expectedError := customErrors.TemporaryError{"some temporary error"}
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, expectedError).Times(3)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.EqualError(err, "some temporary error", "GetByID() should return an error when Get() returns an error.")
	s.Equal(Todo{}, actualTodo, "GetByID() should return an empty Todo.")
}

func (s *TodoTestSuite) TestGetByIDCustomErrorCacheActive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	s.mockCache.EXPECT().Get(testTodo.ID).Return(nil, customErrors.TemporaryError{"some temporary error"}).Times(1)
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, customErrors.TemporaryError{"some temporary error"}).Times(3)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.EqualError(err, "some temporary error", "GetByID() should return an error when Get() returns an error.")
	s.Equal(Todo{}, actualTodo, "GetByID() should return an empty Todo.")
}

func (s *TodoTestSuite) TestGetByIDCustomTemporaryErrorCacheActive() {
	s.T().Log("GetByID() should return an error when Get() returns an error.")

	// Activate Cache
	s.todoDAO.CacheActive = true

	// Mock Expectations
	expectedByteArray, err := json.Marshal(testTodo)
	s.NoError(err)
	s.mockCache.EXPECT().Get(testTodo.ID).Return(nil, customErrors.TemporaryError{"some temporary error"}).Return(expectedByteArray, nil)
	s.mockPersistence.EXPECT().Get(testTodo.ID).Return(nil, customErrors.TemporaryError{"some temporary error"}).Times(3)

	actualTodo, err := s.todoDAO.GetByID(testTodo.ID)
	s.Equal(Todo{}, actualTodo, "GetByID() should return an empty Todo.")
}
