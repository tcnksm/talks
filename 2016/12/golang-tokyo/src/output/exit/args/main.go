package main

import (
	"io"
	"os"
)

// START OMIT
func main() {

	os.Exit(Run(os.Args))
}

type CLI struct {
	out io.Writer
	err io.Writer
}

func Run(args []string) int {
	if err := process1(); err != nil {
		return 1
	}

	if err := process2(); err != nil {
		return 1
	}

	return 0
}

// END OMIT

func process1() error {
	return nil
}

func process2() error {
	return nil
}
