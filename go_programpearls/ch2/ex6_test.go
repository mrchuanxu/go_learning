package ch2_test

import (
	"testing"
)

func BenchmarkThreeTwo(b *testing.B) {
	var a, c int32
	a = 42
	c = 71
	for i := 0; i < b.N; i++ {
		var sum int32
		for j := 0; j < 1000; j++ {
			sum += (a + c)
		}
		b.Log(sum)
	}
}

func BenchmarkSixFour(b *testing.B) {
	var a, c int64
	a = 42
	c = 71
	for i := 0; i < b.N; i++ {
		var sum int64
		for j := 0; j < 1000; j++ {
			sum += (a + c)
		}
		b.Log(sum)
	}
}
