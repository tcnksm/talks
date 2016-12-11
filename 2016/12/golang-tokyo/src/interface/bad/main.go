package main

import "log"

func main() {
	redis := &Redis{
		addr: ":6379",
	}

	name, err := process(redis)
	if err != nil {
		log.Fatal(err)
	}

	_ = name
}

func process(redis *Redis) (string, error) {
	return redis.Get("name")
}

type Redis struct {
	addr string
}

func (r *Redis) Get(key string) (string, error) {
	return "", nil
}
