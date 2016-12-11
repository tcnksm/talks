package main

import "os"

// START OMIT
func main() {
	os.Exit(Run())
}

func Run() int {
	if err := process1(); err != nil {
		return 1 // 🙆
	}

	if err := process2(); err != nil {
		return 1 // 🙆
	}

	return 0 // 🙆
}

// END OMIT

func process1() error {
	return nil
}

func process2() error {
	return nil
}
