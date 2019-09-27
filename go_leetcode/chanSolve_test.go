package chanSolve

import (
	"testing"
	"sync"
	"runtime"
	"fmt"
)

// 交替打印字母与数字
// 控制协程交替打印
const (
	N = 28
)

func TestChanSolve(t *testing.T){
	runtime.GOMAXPROCS(1) // 用一个CPU内核

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i:=0;i<N;i++{
		go func(inum int){
			defer wg.Done()
			fmt.Printf("%d %d",inum+1,inum+2)
			runtime.Gosched() // 让出时间片
		}(i)
		go func(inum int){
			defer wg.Done()
			if inum == 26{
				runtime.Goexit()
			}
			fmt.Printf("%c %c",'A'+inum,'A'+inum+1)
		}(i)
		i++ // 输出了两个 所以计数就+1啦
	}
	wg.Wait()
}