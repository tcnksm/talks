package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	// START OMIT
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set hooks
	trace := httptrace.ClientTrace{
		GetConn: func(h string) {
			log.Println("GetConn:", h)
		},
	}

	ctx := httptrace.WithClientTrace(context.Background(), &trace)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, _ := client.Do(req)
	// END OMIT

	_ = res
}
