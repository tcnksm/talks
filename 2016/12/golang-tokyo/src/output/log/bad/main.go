package main

import "log"

type Foo struct {
}

func (f *Foo) Process() error {
	log.Printf("Make dependencies explicit")
	// ...
	return nil
}
