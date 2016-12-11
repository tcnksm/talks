package main

import "testing"

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
		_ = tc

		// sc := scanner.New(tc.Input)
		// parser := New(sc)

		// program := parser.ParserProgram()
		// checkErrors(t, parser.Errors())

		// if got := len(program.Stmts); got != 1 {
		// 	t.Fatalf("program.Statements contains %d statements, want 1", got)
		// }

		// stmt := program.Stmts[0]

		// val := stmt.(*ast.LetStmt).Value
		// testLiteralExpression(t, val, tc.Value)
	}
}
