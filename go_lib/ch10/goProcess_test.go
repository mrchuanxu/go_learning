package ch10

import (
	"testing"
	// "io/ioutil"
	"os/exec"
	"fmt"
)
// 提供了创建进程的方法

// go语言使用clone 也是很直接
// go 没有直接提供fork系统调用的封装，而是将fork和execve合二为一，提供了syscall.ForkExec
// 如果想只调用fork 通过syscall.Syscall(syscall.SYS_FORK,0,0,0)实现

func TestCmd(t *testing.T){
	dateCmd := exec.Command("date")
	dateOut,err := dateCmd.Output()
	if err!=nil{
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOut))

	// 从外部进程stdin输入数据并从stdout收集结果
	grepCmd := exec.Command("grep","hello")
	// 明确获取输入/输出管道
	
}
