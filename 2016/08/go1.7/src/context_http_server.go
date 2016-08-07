package main

import "net/http"

// START OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// ...
}

// END OMIT
