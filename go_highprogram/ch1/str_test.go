package ch1

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// 字符串 一个不可改变的字节序列 底层数据结构 也是和切片类型，包含一个指向数据的地址 一个长度

type IStringHeader struct {
	Data uintptr
	Len  int
}

// var data = [...]byte{'h','h'} <=> 等价str := "hh"

//  切片 go语言的特色
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func TestStr(t *testing.T) {
	// str := "helllo"
	// str1 := str
	// strings.ToUpper(str)
	// t.Log(str1)
	// s1 := "helloworld"[:5]
	// ls1 := (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len
	// t.Log(ls1)
	// i := 2 | 3
	// t.Log(i)
	// a := []byte{'1', '2', '3', '4'}
	// p := byteToString(a)
	// a = a[:copy(a, a[1:])] // a是已经被改变了
	// t.Log(a)
	// t.Log(p)
	// a = append(a[:2], append([]byte{'1', '2', '3', '4'}, 'a', 'b', 'c')...)

	// t.Log(a)
	// copy(a[2:], a[1:])
	// a[1] = 'd'
	// t.Log(a)

	s1 := "bbz"
	s2 := "adc"
	if s1 > s2 {
		t.Log("sure")
	}
	// b1 := []byte{'h', 'h'}
	fmt.Fprintln(&upwriter{os.Stdout}, "hi world")
}

type upwriter struct {
	io.Writer
}

func (up *upwriter) Write(b []byte) (int, error) {
	return up.Writer.Write(bytes.ToUpper(b))
}

func byteToString(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0])) // 指向data首位指针
	hdr.Len = len(data)
	return
}

// context同步原语

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
