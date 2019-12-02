package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		isHasHTTP := strings.HasPrefix(url, "http://")
		if !isHasHTTP {
			url = "http://" + url
		}
		resp, err := http.Get(url) // 创建http请求的函数
		// resp结构体得到访问的请求结果
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body) // resp的body字段包含一个可读的服务器响应流 iotuil.ReadAll函数从response读取全部内容，将其结果保存在变量b中
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 关闭resp的Body流 防止资源泄漏？
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d %s", b, resp.Status)
	}
}
