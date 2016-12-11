package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if !AskDeploy(os.Stdin) {
		log.Fatal("Cancel deploy")
	}

	log.Println("Deploy!")
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
