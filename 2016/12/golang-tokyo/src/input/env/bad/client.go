package main

import (
	"fmt"
	"os"
)

func NewClient() (*Client, error) {
	token := os.Getenv("GT_TOKEN") // 😇
	if len(token) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	return &Client{
		Token: token,
	}, nil
}

type Client struct {
	Token string
}
