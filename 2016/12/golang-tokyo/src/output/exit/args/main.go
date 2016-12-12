package main

import (
	"io"
	"os"
)

// START OMIT
func main() {
	cli := &CLI{
		out: os.Stdout,
		err: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}

type CLI struct {
	out io.Writer
	err io.Writer
}

func (c *CLI) Run(args []string) int {

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
