package ch5_test

import "testing"

func squares() func() int {
	var i int
	return func() int {
		i++
		return i * i
	}
}

// TestAnoymous 测试匿名函数
func TestAnoymous(t *testing.T) {
	f := squares()
	for i := 0; i < 5; i++ {
		t.Log(f())
	}
}
