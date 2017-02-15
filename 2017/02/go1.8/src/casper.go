package main

import (
	"log"
	"net/http"

	casper "github.com/tcnksm/go-casper"
)

func main() {
	// START OMIT
	pusher := casper.New(1<<6, 10)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := pusher.Push(w, r, []string{"/main.js"}, nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	})
	// END OMIT
}
