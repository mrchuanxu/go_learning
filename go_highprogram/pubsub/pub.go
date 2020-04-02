package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者 一个管道
	topicFunc  func(v interface{}) bool // 主题是一个过滤器
)

type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列缓存大小
	timeout     time.Duration            // 超市时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

func NewPublisher(pubtimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     pubtimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}
