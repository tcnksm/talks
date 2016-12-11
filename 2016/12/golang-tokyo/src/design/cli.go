package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const defaultConfigPath = "./golang-tokyo.json"

type CLI struct {
	outStream io.Writer
	inStream  io.Reader
}

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		inStream:  os.Stdin,
	}

	os.Exit(Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("golang-tokyo", flag.ContinueOnError)
	cfgPath := flags.String("config", defaultConfigPath, "")

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	if AskDeploy(c.inStream) {
		//
	}

	return 0
}

// AskDeploy askes user ok to deploy or not
func AskDeploy(r io.Reader) bool {

	fmt.Println("Do you want to deploy? [y/N]")
	reader := bufio.NewReader(r)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false
	}

	ans := strings.TrimSuffix(line, "\n")
	if ans == "y" || ans == "Y" {
		return true
	}

	return false
}
