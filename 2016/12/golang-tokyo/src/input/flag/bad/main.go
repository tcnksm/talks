package main

import (
	"flag"
	"fmt"
)

func main() {
	var print = flag.Bool("print", false, "")
	flag.Parse()

	if *print {
		fmt.Println("golang tokyo!")
	}
}
