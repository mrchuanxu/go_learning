package ch2_test

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var total uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func TestHighRun(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	t.Log(total)
}

// 同步原语
func TestSync(t *testing.T) {
	done := make(chan int)

	go func() {
		println("hello world")
		done <- 1
	}()
	<-done
}

var bdone = make(chan bool)
var msg string

func aGo() {
	msg = "hi trans"
	<-bdone
}

func TestBdone(t *testing.T) {
	go aGo()
	bdone <- true
	println("!", msg)
}

func TestPrint(t *testing.T) {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func(num int) {
			t.Log("hello trans")
			done <- num
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		t.Log(<-done)
	}
}

func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		t.Logf("%v %v\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime)
	}
	cancel()
}

func gA(a <-chan int) {
	for {
		select {
		case val := <-a:
			fmt.Println(val)
		}
	}
}

func TestGA(t *testing.T) {
	ch := make(chan int)
	ch <- 3
	// go gA(ch)
	// ch <- 5
}
