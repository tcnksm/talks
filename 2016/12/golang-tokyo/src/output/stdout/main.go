package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	calc := &Calculator{
		output: os.Stdout,
	}

	calc.DisplayAdd(1, 2)
}

type Calculator struct {
	output io.Writer
}

func (c *Calculator) DisplayAdd(i, j int) {
	fmt.Fprintf(c.output, "%d + %d = %d\n", i, j, i+j)
}
