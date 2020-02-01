package ch3_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func RandomInt(m, n int) {
	remaining := n
	selectNum := m
	for i := 0; i < n; i++ {
		if rand.Int()%remaining < selectNum {
			fmt.Println(i)
			selectNum--
		}
		remaining--
	}
}

func TestRandomInt(t *testing.T) {
	RandomInt(3, 100)
}
