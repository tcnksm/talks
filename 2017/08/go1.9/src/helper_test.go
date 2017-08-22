package main

import "testing"

// START OMIT
func prepare(t *testing.T) {
	t.Helper()
	t.Fatal("failed to prepare") // line: 6
}

func TestDo(t *testing.T) {
	prepare(t) // line: 11
}

// END OMIT
