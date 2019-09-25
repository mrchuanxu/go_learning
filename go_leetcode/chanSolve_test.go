import (
	"sync"
	"testing"
	"runtime"
	"fmt"

)

func TestChanSolve(t *testing.T){
	runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}

	wg.Add(N)

	for i:=0;i<N;i++{
		go func(inum int){
			defer wg.Done()
			t.Logf("%d %d",inum+1,inum+2)
			runtime.Gosched() // 让出CPU时间片
		}(i)
		go func(inum int){
			defer wg.Done()
			if inum==26{
				runtime.Goexit()
			}
			t.Logf("%c %c",'A'+inum,'A'+inum+1)
		}(i)
		i++
	}
	wg.Wait()
}