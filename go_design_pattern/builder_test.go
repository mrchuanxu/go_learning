package go_design_pattern_test

import (
	"fmt"
	"testing"
)

type buidler struct {
	Name string
	Age  string
	High string
}

func (b *buidler) SetName(name string) *buidler {
	b.Name = name
	return b
}

func (b *buidler) SetAge(age string) *buidler {
	b.Age = age
	return b
}

func (b *buidler) SetHigh(high string) *buidler {
	b.High = high
	return b
}

func Test_builder(t *testing.T) {
	b := new(buidler).SetAge("23").SetHigh("72").SetName("trans")
	fmt.Printf("[%+v]", b)
}
