package gorountine

import (
	"testing"
	//"time"
	"sync"
)

func TestRountine(t *testing.T){
	conter := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i:=0;i<5000;i++{
		wg.Add(1)
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			conter++
			wg.Done()
		}()
	}
	wg.Wait()
	// time.Sleep(time.Millisecond*50)
	t.Log(conter)
}

