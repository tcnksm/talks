package main

import (
	"testing"

	"github.com/hashicorp/hcl/hcl/scanner"
)

func TestParseLetStatement(t *testing.T) {
	cases := []struct {
		Input      string
		Identifier string
		Value      interface{}
	}{
		{`let x = 5;`, "x", 5},
		{`let y = 10;`, "y", 10},
		{`let foobar = y;`, "foobar", "y"},
	}

	for _, tc := range cases {
		sc := scanner.New(tc.Input)
		parser := New(sc)
		program := parser.ParserProgram()
		if got := len(program.Stmts); got != 1 {
			t.Fatalf("program.Statements contains %d statements, want 1", got)
		}
		// ...
	}
}
