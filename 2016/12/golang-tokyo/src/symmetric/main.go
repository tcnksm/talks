package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Name string
}

func GetUser(domain string, id int, testUpdate bool) (*User, error) {
	res, err := http.Get(fmt.Sprintf("%s/users/%d", domain, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if testUpdate {
		f, _ := os.Create(fmt.Sprintf("./testdata/%d", id))
		res.Body = ioutil.NopCloser(io.TeeReader(res.Body, f))
	}

	user := &User{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}
