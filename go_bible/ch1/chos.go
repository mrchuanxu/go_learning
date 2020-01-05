package main

import "os"

import "fmt"

// os以跨平台的方式，提供了一些与操作系统交互的函数和变量
// os.Argss 是一个字符串切片
// s[m:n] 获取子序列 左闭右开形式 区间包含第一个索引元素，不包含最后一个 为了逻辑简单
// talk is cheap
func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	a := 2.51
	fmt.Println(fmt.Sprintf("%.2f", a))
}
