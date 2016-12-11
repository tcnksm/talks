package main

import "testing"

func TestGetUser(t *testing.T) {
	want := "deeeet"
	user, err := GetUser("https://realserver.com", 1)
	if err != nil {
		t.Fatal(err)
	}

	if got := user.Name; got != want {
		t.Fatalf("GetUser = %s, want %s", got, want)
	}
}
