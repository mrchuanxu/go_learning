package interface_test

import (
	"testing"
	"fmt"
	"strconv"
)

type Shape interface{
    Area() float64
    Perimeter() float64
    Circle(width ...float64) float64
}

type Soga interface{
	myTime() int
}

// 建一个Rect
type Rect struct{
	width float64
	height float64
}

func (r Rect)Area()float64{
	return r.width * r.height
}

func (r Rect)Perimeter()float64{
	return 2*(r.width+r.height)
}

func (r Rect)Circle(width ...float64)float64{
	if len(width)==0{
		return 0
	}
	isum := 0.0
	for _,iwidth := range(width){
		isum +=iwidth
	}

	return isum
}

func (r Rect) Time()float64{
	return 3.14
}

func (r Rect) myTime()int{
	return 233
}
type MyString string

func explain(i interface{}){
	
	// 要是这里使用i 就需要断言判断了
	switch val :=i.(type){
	case string:
		fmt.Println(val)
	case int:
		val = val+1
		fmt.Println(val)
	default:
		fmt.Printf("%v,%T",i,i)
	}
}

type Stringer interface{
	String() string
}

type Binary uint64

func(i Binary) String() string{
	return strconv(i.Get(),2)
}
func (i Binary)Get() uint64{
	return uint64(i)
}

func TestInterface(t *testing.T){
	// var s Shape = Rect{3.14,4}
	// c := s.(Rect) // 这里返回的是Rect类型 这里提取了
	// fmt.Println(c.myTime())
	//explain(1)
	b := Binary{}
	s := Stringer(b)
	t.Log(s.String())
}
