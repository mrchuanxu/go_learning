package ch_test

import (
	"fmt"
	"image/color"
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	defer func() {
		if p := recover(); p != 1 {
			fmt.Printf("internal error:%v", p)
		} else {
			panic(p)
		}
	}()
	time.AfterFunc(1*time.Second, func() {
		t.Log(Distance(10))
	})
	a := 2
	if a == 1 {
		panic(a)
	} else {
		panic(a)
	}
}

func Distance(n int) int {
	return n
}

type Point struct{ x, y float64 }

type ColorPoint struct {
	Point
	Color color.RGBA64
}

func TestPoint(t *testing.T) {
	var cp ColorPoint
	cp.Point.x = 1
}
