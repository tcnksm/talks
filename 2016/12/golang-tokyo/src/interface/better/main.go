package main

import "log"

func main() {

	db := NewRedis(":6379")
	name, err := process(db)
	if err != nil {
		log.Fatal(err)
	}

	_ = name
}

func process(db DB) (string, error) {
	return db.Get("name")
}

type DB interface {
	Get(key string) (string, error)
}

func NewRedis(addr string) DB {
	return &Redis{
		addr: addr,
	}
}

type Redis struct {
	addr string
}

func (r *Redis) Get(key string) (string, error) {
	return "", nil
}
