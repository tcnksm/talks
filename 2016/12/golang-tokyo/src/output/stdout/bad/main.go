package main

import "fmt"

func main() {
	calc := &Calculator{}

	calc.DisplayAdd(1, 2)
}

// START OMIT
type Calculator struct {
}

func (c *Calculator) DisplayAdd(i, j int) {
	fmt.Printf("%d + %d = %d\n", i, j, i+j)
}

// END OMIT

// // Printf formats according to a format specifier and writes to standard output.
// // It returns the number of bytes written and any write error encountered.
// func Printf(format string, a ...interface{}) (n int, err error) {
// 	return Fprintf(os.Stdout, format, a...)
// }
