package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func TestHandler(t *testing.T) {
	cases := []struct {
		Name   string
		Status int
		Body   string
	}{
		{"deeeet", http.StatusOK, "Hello, deeeet\n"},
		{"", http.StatusBadRequest, ""},
	}

	for _, tc := range cases {
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://hello.com?name=%s", tc.Name), nil)
		w := httptest.NewRecorder()
		helloHandler(w, req)

		if got, want := w.Code, tc.Status; got != want {
			t.Fatalf("StatusCode=%d, want=%d", got, want)
		}

		if got, want := w.Body.String(), tc.Body; got != want {
			t.Fatalf("Body=%q, want=%q", got, want)
		}
	}
}
