package main

import (
	"os"
	"testing"
)

func setEnv(k, v string) func() {
	prev := os.Getenv(k)
	os.Setenv(k, v)
	return func() {
		os.Setenv(k, prev)
	}
}

func TestStatusHandler(t *testing.T) {
	reset := setEnv(EnvStatus, "Maintainance!")
	defer reset()

	// test
}
