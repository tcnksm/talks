package main

import "os"

func main() {
	if err := process1(); err != nil {
		os.Exit(1)
	}
	if err := process2(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func process1() error {
	return nil
}

func process2() error {
	return nil
}
