package main

import (
	"log"
	"net/http"
)

func main() {
	// START OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pusher, ok := w.(http.Pusher)
		if ok {
			if err := pusher.Push("/main.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
	})
	// END OMIT
}
