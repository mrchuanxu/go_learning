### golang 的文件锁操作
多个goroutine操作同一个文件，需要一个文件锁🔒Flock。文字也是显而易见。<br>
Flock是对整个文件建议性锁(不强求goroutine遵守)，如果一个goroutine在文件上获取锁🔒，那么其他的goroutine是可以知道的。默认情况下，当一个goroutine将文件锁住了，另外一个goroutine可以直接操作被锁住的文件，原因在于Flock只是用于检测文件是否被加锁，针对文件已经被加锁，另一个goroutine写入数据的情况下，内核不会阻止这个goroutine的写入操作，也就是建议性锁的内核处理策略。<br>
##### 函数
```go
import "syscall"
func Flock(fd int,how int)(err error)
```
Flock位于syscall包中，fd参数指代文件描述符，how参数指代锁的操作类型。<br>
[传送门](https://reading.developerlearning.cn/articles/2018-11-11-golang-file-lock/)

* Flock是建议性锁，使用的时候需要指定how参数，否则容易出现多个goroutine共用文件的问题
* how参数指定LOCK_NB之后，goroutine遇到已加锁的Flock，不会阻塞，而是直接返回错误
