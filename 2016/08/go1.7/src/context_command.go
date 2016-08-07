package main

import (
	"log"
	"os/exec"
)

func main() {
	// START OMIT
	cmd := exec.CommandContext(ctx, "awesome_command")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
