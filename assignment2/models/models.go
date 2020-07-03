package models

type Persistence interface {
	Put(key string, value []byte) error // Put() updates or replaces resources in the cache based on the existence of said resource
	Get(key string) ([]byte, error)     // Get() retrieves a record by the specified key
}

type CacheInterface interface {
	Persistence
	Flush(key string)
}
