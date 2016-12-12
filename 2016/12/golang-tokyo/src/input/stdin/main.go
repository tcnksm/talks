package main

import (
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

	var line string
	if _, err := fmt.Fscanln(r, &line); err != nil {
		return false
	}

	ans := strings.TrimSuffix(line, "\n")
	if ans == "y" || ans == "Y" {
		return true
	}

	return false
}
