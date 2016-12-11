package main

import "os"

func main() {
	os.Exit(Run(os.Args))
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

func process1() error {
	return nil
}

func process2() error {
	return nil
}
