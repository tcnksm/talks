package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	if !AskDeploy() {
		log.Fatal("Cancel deploy")
	}

	log.Println("Deploy!")
}

// Yes/Noでユーザに回答を求める
func AskDeploy() bool {

	fmt.Println("Do you want to deploy? [y/N]")

	var line string
	if _, err := fmt.Scanln(&line); err != nil {
		return false
	}

	ans := strings.TrimSuffix(line, "\n")
	if ans == "y" || ans == "Y" {
		return true
	}

	return false
}
