package main

import "testing"

func BenchmarkFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
