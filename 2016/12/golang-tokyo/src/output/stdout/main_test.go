package main

import (
	"bytes"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		Input1, Input2 int
		Want           string
	}{
		{1, 2, "1 + 2 = 3\n"},
		{100, 100, "100 + 100 = 200\n"},
	}

	for _, tc := range cases {
		var buf bytes.Buffer
		calc := &Calculator{
			output: &buf,
		}

		calc.DisplayAdd(tc.Input1, tc.Input2)
		if got := buf.String(); got != tc.Want {
			t.Fatalf("DisplayAdd shows %q, want %q", got, tc.Want)
		}
	}
}
