package ch1_test

import (
	"fmt"
	"testing"
)

// 匿名函数
// 具名函数需要先声明才能使用
// go语言程序先初始化 而后再执行

func TestFunc(t *testing.T) {
	t.Log("hi func")
}

type AA interface {
	name() string
	age() string
}

type BB interface {
	AA
	gender() string
}

type Runner interface {
	Run()
	Say()
}

type Person struct {
	Name string
}

func (p Person) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *Person) Say() {
	fmt.Printf("hello, %s", p.Name)
}

func TestInter(t *testing.T) {
	var r Runner
	r = &Person{Name: "song_chh"}
	r.Run()
	r.Say()
}
