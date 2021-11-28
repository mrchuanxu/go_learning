package go_design_pattern_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

// 将请求的发送和接收解耦，让多个接收对象都有机会处理这个请求。 将这些接收对象串成一条链，并沿着这条链传递这个请求，直到链上的某个接收对象能够处理它。

type IHandler interface {
	Process(ctx context.Context, tag string) bool
}

type Processer struct {
	Tag string
	IHandler
}

type HandlerChain struct {
	Processers []*Processer // 专门用于存储职责链下的处理
}

func (h *HandlerChain) addProcess(processer *Processer) {
	h.Processers = append(h.Processers, processer)
}

func (p *Processer) Process(ctx context.Context, tag string) bool {
	if p.Tag == tag {
		return p.IHandler.Process(ctx, tag)
	}
	return false
}

func (h *HandlerChain) handlerProcess(ctx context.Context, tag string) error {
	for i := range h.Processers {
		if h.Processers[i].Process(ctx, tag) {
			return nil
		}
	}
	return errors.New("no handler is done")
}

type HandlerOne struct {
}

func (h *HandlerOne) Process(ctx context.Context, tag string) bool {
	fmt.Printf("handlerOne tag:[%s]", tag)
	return true
}

type HandlerTwo struct {
}

func (h *HandlerTwo) Process(ctx context.Context, tag string) bool {
	fmt.Printf("handlerOne tag:[%s]", tag)
	return true
}

func Test_handlerChain(t *testing.T) {
	chain := &HandlerChain{}
	p1 := &Processer{"s", &HandlerOne{}}
	chain.addProcess(p1)
	p2 := &Processer{"b", &HandlerOne{}}
	chain.addProcess(p2)
	chain.handlerProcess(context.Background(), "b")
}

// 常用的过滤器、拦截器的实现
// 定义一个过滤器实现，过滤器自动处理
// 让框架的使用者在不需要修改框架源码的情况下，添加新的过滤拦截的功能。 对扩展开放，对修改关闭。
var pos int
var n = 0
var Filters = []func(string, string) error{}

func doFilter(request, resp string) error {
	if pos < n {
		filter := Filters[pos]
		// 拦截客户端请求
		fmt.Println("request filter")
		filter(request, resp)
		// 拦截发给客户端的响应
		fmt.Println("response filter")
	} else {
		fmt.Println("fiilter Done")
	}
	return nil
}
