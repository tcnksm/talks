package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		Command string
		Status  int
	}{
		{
			"cool-command -config test.json arg1 arg2",
			0,
		},

		{
			"cool-command -v",
			1,
		},
	}

	for _, tc := range cases {
		args := strings.Split(tc.Command, " ")
		if got := Run(args); got != tc.Status {
			t.Fatalf("Run exit %d, want = %d", got, tc.Status)
		}
	}
}
