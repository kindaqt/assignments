package models

import (
	"github.com/jinzhu/gorm"
)

///////////////////////////////////
// Persister
////////////////////////////////

type psqlStore struct {
	Config
	DB *gorm.DB
}

type Config struct {
	Driver     string
	ConnString string
}

func NewPsqlStore(driver string, connString string) (Persistence, error) {
	store := &psqlStore{
		Config: Config{
			Driver:     driver,
			ConnString: connString,
		},
	}
	if err := store.Connect(); err != nil {
		return nil, err
	}
	return store, nil
}

////////////////////////////
// Database
//////////////////////////

func (p *psqlStore) Connect() error {
	// Open DB
	db, err := gorm.Open(p.Config.Driver, p.Config.ConnString)
	if err != nil {
		return err
	}

	if err := db.DB().Ping(); err != nil {
		return err
	}

	// Attach DB to persister
	p.DB = db

	return nil
}

func (p *psqlStore) Close() error {
	if err := p.DB.Close(); err != nil {
		return err
	}
	return nil
}

/////////////////////////////////
// Methods
///////////////////////////////

// TodoGormModel is a Model for Gorm
type TodoGormModel struct {
	Key   string `gorm:"column:key;primary_key"`
	Value []byte `gorm:"column:value"`
}

// Put() puts a record in the todos table
func (p *psqlStore) Put(key string, value []byte) error {
	return p.DB.Save(&TodoGormModel{key, value}).Table("todos").Error
}

// Get() gets a record from the todos table
func (p *psqlStore) Get(key string) ([]byte, error) {
	var record TodoGormModel
	if err := p.DB.Select("value").Where("key = ?", key).First(&record).Error; err != nil {
		return nil, err
	}
	return record.Value, nil
}
