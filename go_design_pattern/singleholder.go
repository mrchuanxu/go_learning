package go_design_pattern

import "sync"

// 考虑实例的线程安全的问题
// 单例模式能解决

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
