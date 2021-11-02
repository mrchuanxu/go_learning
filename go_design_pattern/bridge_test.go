package go_design_pattern_test

import (
	"fmt"
	"testing"
)

// 将抽象和实现解耦，让它们可以独立变化
// 抽象
type PrinterIFace interface {
	PrintFile()
}

type A struct {
	printer PrinterIFace
}

func (a *A) Print() {
	a.printer.PrintFile()
}
func (a *A) setPrinter(p PrinterIFace) {
	a.printer = p
}

type B struct {
	printer PrinterIFace
}

func (b *B) Print() {
	b.printer.PrintFile()
}
func (b *B) setPrinter(p PrinterIFace) {
	b.printer = p
}

// 实现
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("hello epson")
}

func Test_Print(t *testing.T) {
	b := &B{}
	b.setPrinter(&Epson{})
	b.Print()
}
