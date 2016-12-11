package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

func GetUser(domain string, id int) (*User, error) {
	res, err := http.Get(fmt.Sprintf("%s/users/%d", domain, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	user := &User{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
