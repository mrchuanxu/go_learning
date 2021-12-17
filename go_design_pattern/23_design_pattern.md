# 23种设计模式随想
思从深而行从简。<br>

花了三个月时间，读完极客时间设计模式之美，从设计原则到设计模式，到场景应用，作者句句戒点睛之笔，深究其言，亦是回味无穷，思与行之间，代码设计之美体现淋漓尽致。方寸把握亦是老道，经验之足，不是我现在能比拟的。现就设计模式之美的学习笔记以及心得作一总结以便日后持续参考回味，更新可能会在一周内更新完毕。

## 鸟瞰
23种设计模式分为创建型、结构型、行为型三种类型，其实大多数源码也是围绕这三种类型展开设计。其中
* 创建型包括单例模式、工厂模式、建造者模式、原型模式；
* 结构型包括代理模式、适配器模式、桥接模式、装饰器模式、门面模式、组合模式、享元模式；
* 行为型包括职责链模式、中介模式、访问者模式、模版模式、策略模式、状态模式、迭代器模式、备忘录模式、命令模式、解释器模式。

设计模式可以灵活运用在各种场景，各种场景也不止可以用一种模式进行设计，可以揉和设计。

## 讲述结构
通篇将以原理，实现，用例场景结合实战，优点与缺点四个标题来讲述每一个设计模式。
### 单例模式
一个类只允许创建一个对象或者实例，那这个类就是一个单例类。为什么要使用单例模式？ 因为单例模式能处理资源访问冲突，并且保持全局唯一，无论是线程上，进程上还是分布式上。 单例的实现方式？
#### 如何用go实现一个日志打印功能
先来看看一个日志打印功能
```go
// 一个logger打印日志就出来了
func logger(ctx context.Context, formatter string, args ...interface{}) {
	// trace ctx
	fd, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = fd.WriteString(fmt.Sprintf(formatter, args...))
	if err != nil {
		panic(err)
	}
	fd.Sync()
}
```
这段代码有什么问题呢？ 假设多协程一起执行写日志，每个协程执行不共享一个文件描述符，就会导致覆盖写。最后输出的只有一行。那如何避免覆盖写？ <br>
方式一: 加`OS_APPEND`
```go
// 一个logger打印日志就出来了
func logger(ctx context.Context, formatter string, args ...interface{}) {
// trace ctx
    fd, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        panic(err)
    }
    defer fd.Close()
    _, err = fd.WriteString(fmt.Sprintf(formatter, args...))
    if err != nil {
        panic(err)
    }
    fd.Sync()
}
```
添加`os.O_APPEND`就不会覆盖写的情况。为什么呢？这是在系统级别上控制，当打开文件设置了追加写，那么内核就会共享文件写入游标，保证内容不会被覆盖。<br>

方式二： 共享一个单例<br>
```go
func logger(fd *os.File, formatter string, args ...interface{}) {
	_, err := fd.WriteString(fmt.Sprintf(formatter, args...))
	if err != nil {
		panic(err)
	}
	fd.Sync()
}
```
在这里，整个函数共享同一个文件描述符，文件描述符需要的是全局唯一，就是整个进程的唯一，在这个进程中，无论协程还是线程对文件的写入都不会存在覆盖写。<br>
实现一个单例，有很多种方式，可以延迟加载，可以提前加载，也可以双重检测等。但是最后的目的只有一个就是这个单例必须符合线程/协程安全，支持高并发使用。<br>
go实现单例的方式
```go
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
```
#### 适用场景
单例适用什么场景？ 单例模式适用于需要频繁被创建使用并销毁的实例。例如配置文件访问，数据库/redis连接池，自增雪花id等。<br>
配置文件访问是因为系统内部进行某些特性功能如邮件短信，某些链接地址信息等读取，只需要新建一个实例即可。<br>
数据库链接频繁创建和销毁很耗费性能，在业务高频应用的情况下，池化是很重要的作用，池化用到的就是单例创建模式，用完链接就放回去。<br>
#### 优点与缺点判断
优点：
1. 提供了一个可控的访问实例。
2. 节约系统资源。

缺点：
1. 对OOP不友好，无法扩展。
2. 对垃圾回收机制不友好，过长时间不用，GC标记为可回收对象，如果此时使用就会发生未知错误。
3. 可测试性差，不能mock
