package main

import "testing"

func Add(i, j int) int {
	return i + j
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Input1, Input2 int
		Want           int
	}{
		{1, 2, 3},
		{100, 100, 200},
		{0, 0, 0},
	}

	for _, tc := range cases {
		if got := Add(tc.Input1, tc.Input2); got != tc.Want {
			t.Errorf("Add=%d, want=%d", got, tc.Want)
		}
	}
}
