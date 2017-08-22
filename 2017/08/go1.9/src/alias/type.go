package main

import (
	"fmt"

	"github.com/tcnksm/talks/2017/08/go1.9/src/alias/api"
	"github.com/tcnksm/talks/2017/08/go1.9/src/alias/util"
)

func main() {
	err := Do()
	_, ok := err.(*api.SpecialError)
	fmt.Println(ok)
}

func Do() error {
	return &util.SpecialError{}
}
