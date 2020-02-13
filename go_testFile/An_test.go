package gofile

import (
	"testing"
)

func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func TestFile(t *testing.T) {
	t.Log(f3())
	t.Log(uint8(')'))
	a := "abcdefg"
	t.Log(a[len(a)-2:])
}
