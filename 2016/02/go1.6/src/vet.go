package main

import "fmt"

func f() string {
	return "Hello"
}

func main() {
	fmt.Printf("%v", f)
	// main.go:10: arg f in printf call is a function value, not a function call
}
