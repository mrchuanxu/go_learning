package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	counts := make(map[string]int) // 传的是一个指针/引用 也是拷贝的指针 指向同一块内存

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts) // 流的模式读取输入
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // 流的模式？文件模式
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; ok {
			fmt.Printf("%s", f.Name())
		}
		counts[input.Text()]++
	}
}
