package ch1

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
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
	Data uintptr // 指向数据地址的指针
	Len  int
	Cap  int
}

func TestStr(t *testing.T) {
	s1 := "bbz"
	s2 := "adc"
	if s1 > s2 {
		t.Log("sure")
	}
	// b1 := []byte{'h', 'h'}
	fmt.Fprintln(&upwriter{os.Stdout}, "hi world")

	// 修改切片
	a := []int{1, 2, 3, 4, 4, 5, 6}
	y := a[1:3:5]
	t.Log(cap(y))
	t.Log(len(y))
	t.Log(y[1], y[0])
	y[1] = 8
	t.Log(a)
	a = append([]int{-3, -2, -1}, a...)
	t.Log(a)
	copy(a[0:], a[1:])
	t.Log(a)
	t.Log(cap(a))
	a = append(a[:0], a[2:]...)
	t.Log(a)
	t.Log(cap(a))
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

func TestSlice(t *testing.T) {
	sl := []string{
		"a",
		"b",
		"c",
	}

	fmt.Printf("%v, %p\n", sl, sl)
	changeSlice(sl)
	fmt.Printf("%v, %p\n", sl, sl)
	changeSlice1(sl)
	fmt.Printf("%v, %p\n", sl, &sl)
}

func changeSlice(sl []string) {
	fmt.Printf("%v, %p\n", sl, sl)
	sl[0] = "aa"

	//sl = append(sl, "d")
	fmt.Printf("%v, %p\n", sl, sl)
}

func changeSlice1(sl []string) {
	fmt.Printf("%v, %p\n", sl, &sl)
}

func TestMap(t *testing.T) {
	aMap := make(map[string]int)
	t.Logf("%p", aMap)
}

func TestInt(t *testing.T) {
	bigInt := ^uint(0)
	t.Log(bigInt)
	t.Log(int(bigInt >> 1))
}

func TestMutile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go work(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		ctx.Err()
	default:
		fmt.Println("Hello lijia!")
	}
}
