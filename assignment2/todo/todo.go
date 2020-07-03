package todo

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kindaqt/assignment2/models"
	"github.com/kindaqt/assignment2/utils/retry"
)

// Todo object
type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// TodoDAO is the Todo Data Access Object Interface
type TodoDAO interface {
	Save(t Todo) error
	GetByID(id string) (Todo, error)
}

// TodoDAOPersister persists resources
type TodoDAOPersister struct {
	DataStore   models.Persistence
	Cache       models.CacheInterface
	CacheActive bool // cache data when true
}

// NewTodoDAO returns a new TodoDAO
func NewTodoDAO(persister models.Persistence, cacheActive bool, cache models.CacheInterface) TodoDAO {
	log.Println("Creating New TodoDAO")

	dao := &TodoDAOPersister{
		DataStore:   persister,
		CacheActive: cacheActive,
		Cache:       cache,
	}
	log.Println("Created New TodoDAO")
	return dao
}

//////////////////////////////
// Methods
////////////////////////////

// Save stores a todo in the repository
func (p *TodoDAOPersister) Save(t Todo) error {
	log.Printf("Saving Todo: %v", t)

	var err error

	// Marshal JSON
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	// Cache data and remove it after successful write
	if p.CacheActive {
		if err := p.Cache.Put(t.ID, b); err != nil {
			log.Println(err)
		} else {
			defer func(err *error) {
				if err == nil {
					p.Cache.Flush(t.ID)
				}
			}(&err)
		}
	}
	// Store Data
	if err := retry.Do(3, time.Duration(time.Millisecond*500), func() error {
		return p.DataStore.Put(t.ID, b)
	}); err != nil {
		return err
	}

	return nil
}

// GetByID returns a todo based on its id
func (p *TodoDAOPersister) GetByID(id string) (Todo, error) {
	log.Printf("Getting Todo by ID: %v", id)
	// Initialize return variable
	var todo Todo
	var todoBytes []byte

	// Get Todo
	if p.CacheActive {
		// Get todo from cache
		if b, err := p.Cache.Get(id); err == nil {
			todoBytes = b
		}
	}
	if err := retry.Do(3, time.Duration(time.Millisecond*400), func() error {
		// Get todo from datastore
		b, err := p.DataStore.Get(id)
		if err == nil {
			todoBytes = b
		}
		return err
	}); err != nil {
		return todo, err
	}

	// Unmarshal Json
	return todo, json.Unmarshal(todoBytes, &todo)
}
