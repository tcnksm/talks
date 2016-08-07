package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

	// START OMIT
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, _ := client.Do(req)

	// END OMIT
}
