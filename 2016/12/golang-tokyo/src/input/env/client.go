package main

import (
	"log"
	"os"
)

func NewClient(token string) (*Client, error) {
	return &Client{
		Token: token,
	}, nil
}

func main() {
	token := os.Getenv("GT_TOKEN") // 🙆
	if len(token) == 0 {
		log.Fatal("missgin token")
	}

	client, err := NewClient(token) // 🙆
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

type Client struct {
	Token string
}
