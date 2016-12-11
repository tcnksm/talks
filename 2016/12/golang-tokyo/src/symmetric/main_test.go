package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	mux.Handle("/users/", http.StripPrefix("/users", http.FileServer(http.Dir("./testdata"))))

	want := "deeeet"
	user, err := GetUser(ts.URL, 1, false)
	if err != nil {
		t.Fatal(err)
	}

	if got := user.Name; got != want {
		t.Fatalf("GetUser = %s, want %s", got, want)
	}
}
