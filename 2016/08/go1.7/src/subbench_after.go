package main

import (
	"fmt"
	"testing"
)

// START OMIT
func BenchmarkFoo(b *testing.B) {
	cases := []struct {
		Base int
	}{
		{Base: 1},
		{Base: 10},
		{Base: 100},
	}

	for _, bc := range cases {
		b.Run(fmt.Sprintf("%d", bc.Base), func(b *testing.B) { benchFoo(b, bc.Base) })
	}
}

// END OMIT
