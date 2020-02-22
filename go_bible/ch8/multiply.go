package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	listener net.Listener
	err      error
)

func init() {
	listener, err = net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("hello internet!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go HandleConn(conn)
	}
}

func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
}
