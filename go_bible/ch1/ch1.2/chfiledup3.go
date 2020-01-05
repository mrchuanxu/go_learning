package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 将数据一口气全部输入数据读到内存中 一次分割为多行
func main3() {
	// ReadFile函数 来自(io/ioutil)包
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
