package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	flags := flag.NewFlagSet("golang-tokyo", flag.ContinueOnError)
	var print = flags.Bool("print", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		log.Fatal(err)
	}

	if *print {
		fmt.Println("golang tokyo!")
	}
}
