package main

import "testing"

// START OMIT
func BenchmarkFoo1(b *testing.B)   { benchFoo(b, 1) }
func BenchmarkFoo10(b *testing.B)  { benchFoo(b, 10) }
func BenchmarkFoo100(b *testing.B) { benchFoo(b, 100) }

func benchFoo(b *testing.B, base int) {
	for i := 0; i < b.N; i++ {
		Foo(base)
	}
}

// END OMIT
