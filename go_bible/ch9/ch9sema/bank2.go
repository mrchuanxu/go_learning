package main

import (
	"fmt"
	"time"
)

// var (
// 	sema    = make(chan struct{}, 1) //  binary semaphore 0 1 信号量
// 	balance int
// )

// // 利用阻塞原理
// func Deposit(amount int) {
// 	sema <- struct{}{}
// 	balance = balance + amount
// 	<-sema
// }

// func Balance() int {
// 	sema <- struct{}{}
// 	b := balance
// 	<-sema
// 	return b
// }

// var mu sync.RWMutex
// var balance int

// func Balance() int {
// 	mu.RLock()
// 	defer mu.RUnlock()
// 	return balance
// }
func main() {
	// ping
	// go func() {
	// 	for {
	// 		select {
	// 		case str := <-pinpon:
	// 			fmt.Println("pong", str)
	// 			pinpon <- "pong"
	// 			time.Sleep(time.Millisecond)
	// 		default:
	// 			pinpon <- "pong"
	// 		}
	// 	}
	// }()

	// pong
	// go func() {
	// 	var pinpon = make(chan string)
	// 	for {
	// 		select {
	// 		case str := <-pinpon:
	// 			fmt.Println("ping", str)
	// 		case pinpon <- "pong":
	// 		default:
	// 		}
	// 	}
	// }()
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
	time.Sleep(time.Second * 10)
}
