package main

import "log"

// START OMIT
func main() {
	redis := &Redis{
		addr: ":6379",
	}

	name, err := process(redis)
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

func process(redis *Redis) (string, error) {
	// ...
}

// END OMIT

type Redis struct {
	addr string
}

func (r *Redis) Get(key string) (string, error) {
	return "", nil
}
