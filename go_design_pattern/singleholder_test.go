package go_design_pattern_test

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

// 考虑实例的线程安全的问题
// 单例模式能解决
// 避免资源访问冲突、表示业务概念上的全局唯一类

type singleTon struct{}

var (
	singleInstance *singleTon
	once           sync.Once
)

// 利用synce的once操作实现单例
// atomic.LoadUint32(&o.done) 源码通过简单的原子操作共享变量，实现单例创建
func GetInstance() *singleTon {
	once.Do(func() {
		singleInstance = &singleTon{}
	})
	return singleInstance
}

// 一个单例模式的打印功能
// 设计的初衷是为了避免打印的时候出现日志文字覆盖

// 一个logger打印日志就出来了
func logger(fd *os.File, formatter string, args ...interface{}) {
	_, err := fd.WriteString(fmt.Sprintf(formatter, args...))
	if err != nil {
		panic(err)
	}
	fd.Sync()
}

func TestLogger(t *testing.T) {
	i := 0
	var wg sync.WaitGroup
	fd, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	for i < 1000 {
		i++
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			logger(fd, "hello %v\n", k)
		}(i)
	}
	wg.Wait()
}
