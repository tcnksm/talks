package main

import "testing"

func TestProcess(t *testing.T) {
	redis := &Redis{
		addr: ":6379",
	}

	want := "golang-tokyo"
	got, err := process(redis)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("process=%s, want = %s", got, want)
	}
}
