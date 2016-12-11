package main

import (
	"io/ioutil"
	"log"
	"testing"
)

var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func TestProcess(t *testing.T) {
	foo := &Foo{
		logger: discardLogger,
	}

	// ...
}
