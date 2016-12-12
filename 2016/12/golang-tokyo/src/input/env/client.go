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
	token := os.Getenv("GT_TOKEN") // mainでローカル変数に代入する
	if len(token) == 0 {
		log.Fatal("missgin token")
	}

	client, err := NewClient(token)
	// ...
}

type Client struct {
	Token string
}
