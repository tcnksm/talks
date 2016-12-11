package main

import (
	"flag"
	"os"
)

const defaultConfigPath = "./golang-tokyo.json"

func main() {
	os.Exit(Run(os.Args))
}

func Run(args []string) int {
	flags := flag.NewFlagSet("golang-tokyo", flag.ContinueOnError)

	var (
		cfgPath = flags.String("config", defaultConfigPath, "")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	if err := process1(); err != nil {
		return 1
	}

	if err := process2(); err != nil {
		return 1
	}

	_ = cfgPath

	return 0
}

func process1() error {
	return nil
}

func process2() error {
	return nil
}
