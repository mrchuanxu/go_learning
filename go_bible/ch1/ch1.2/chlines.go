package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1() {
	counts := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin) // 程序使用短变量声明创建bufio.Scanner类型的变量input
	for inputs.Scan() {                  // 调用Scan() 读入下一行 并移除行末的换行符
		counts[inputs.Text()]++ // Text() 得到读取内容
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			// 类似于C语言的printf函数 对一些表达式格式化输出
			// 以字母f结尾的格式化函数，如log.Printf和fmt.Errorf，都采用fmt.Printf的格式化准则。而以ln结尾的格式化函数，则遵循Println的方式，以跟%v差不多的方式格式化参数，并在最后添加一个换行符。（译注：后缀f指format，ln指line
		}
	}
}
