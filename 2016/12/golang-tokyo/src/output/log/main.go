package main

import (
	"log"
	"os"
)

func main() {
	foo := &Foo{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	foo.Process()
}

type Foo struct {
	logger *log.Logger
}

func (f *Foo) Process() error {
	log.Printf("Make dependencies explicit")
	// ...
	return nil
}
