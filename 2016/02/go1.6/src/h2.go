package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	certFile, _ := filepath.Abs("src/certificate/server.crt")
	keyFile, _ := filepath.Abs("src/certificate/server.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})

	if err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil); err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
