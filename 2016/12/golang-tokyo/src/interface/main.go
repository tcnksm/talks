package main

import "log"

type DB interface {
	Get(key string) (string, error)
}

// START OMIT
func main() {
	db := NewRedis(":6379")
	name, err := process(db)
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

func process(db DB) (string, error) {
	return db.Get("name")
}

func NewRedis(addr string) DB {
	return &Redis{
		addr: addr,
	}
}

// END OMIT
type Redis struct {
	addr string
}

func (r *Redis) Get(key string) (string, error) {
	return "", nil
}
