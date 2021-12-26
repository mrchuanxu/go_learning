package setfinalizer_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

type IApp interface {
	Stop()
}

type App struct {
	running   bool
	stop      chan struct{}
	onStopped func()
}

func (a *App) watch() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-a.stop:
			if a.onStopped != nil {
				a.onStopped()
			}
			return
		case <-ticker.C:
			fmt.Println("not stop")
		}
	}
}

func (a *App) Stop() {
	if !a.running {
		return
	}
	close(a.stop)
}

func New() *App {
	a := &App{
		running: true,
		stop:    make(chan struct{}),
	}
	go a.watch()
	return a
}

func Test_IAPP(t *testing.T) {
	app := NewCache()
	var cnt = 0
	stopped := make(chan struct{})
	app.onStopped = func() {
		cnt++
		close(stopped)
	}
	t.Logf("cnt:%v", cnt)
	app = nil
	runtime.GC()

	select {
	case <-stopped:
	case <-time.After(time.Second * 10):
		t.Fail()
	}
	t.Logf("done cnt:%v", cnt)
}

type Cache = *wrapper

type wrapper struct {
	*cache
}

type cache struct {
	stopped   chan struct{}
	onStopped func()
}

func newCache() *cache {
	return &cache{
		stopped: make(chan struct{}),
	}
}

func NewCache() Cache {
	w := &wrapper{
		cache: newCache(),
	}
	go w.cache.run()
	runtime.SetFinalizer(w, (*wrapper).stop) // 注册执行goroutine最后终止的动作。使得gc回收不可达对象。
	// 当gc检测到unreachable对象有关联的SetFinalizer函数时，会执行关联的SetFinalizer函数， 同时取消关联。 这样当下一次gc的时候，对象重新处于unreachable状态并且没有SetFinalizer关联， 就会被回收。
	return w
}

func (w *wrapper) stop() {
	w.cache.stop()
}
func (c *cache) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.stopped:
			if c.onStopped != nil {
				c.onStopped()
			}
			return
		case <-ticker.C:
			fmt.Println("not stop")
		}
	}
}

func (c *cache) stop() {
	close(c.stopped)
}
