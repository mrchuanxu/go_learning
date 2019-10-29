package gojsonp

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	inum := 42
	setPointer(&inum)
	getPointer(&inum)
}

func setPointer(inum *int) {
	*inum = 12
}

func getPointer(inum *int) {
	fmt.Println(*inum)
}
