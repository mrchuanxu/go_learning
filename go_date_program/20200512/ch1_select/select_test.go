package select_test

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	c, quit := make(chan int), make(chan int)
	go fib(c, quit)
	go func() {
		quit <- 100
		t.Log("quit input")
	}()
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit quit")
			return
		}
	}
}
