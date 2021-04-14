package fib

import "testing"

func BenchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib20(b *testing.B) {
	BenchmarkFib(b, 20)
}
