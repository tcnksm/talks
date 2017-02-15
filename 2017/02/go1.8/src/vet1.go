package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://golang.org")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = res
}
