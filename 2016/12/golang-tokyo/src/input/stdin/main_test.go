package main

import (
	"bytes"
	"testing"
)

func TestAskDeploy(t *testing.T) {
	cases := []struct {
		Input *bytes.Buffer
		Want  bool
	}{
		{bytes.NewBufferString("Y\n"), true},
		{bytes.NewBufferString("n\n"), false},
	}

	for _, tc := range cases {
		if got := AskDeploy(tc.Input); got != tc.Want {
			t.Fatal("AskDeploy %v, want %v", got, tc.Want)
		}
	}
}
