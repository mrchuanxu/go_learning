package go_design_pattern_test

import "testing"

// 在不违背封装原则的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，以便恢复对象先前的状态。

// :list 展示输入， :sudo回撤输入

// 设计一个栈
type Stack struct {
	buffer []string
	len    int // 从0开始算起
}

func (s *Stack) Top() string {
	if len(s.buffer) == 0 {
		return ""
	}
	// 取最后的元素
	return s.buffer[s.len]
}

func (s *Stack) Pop() string {
	if len(s.buffer) == 0 {
		return ""
	}
	ps := s.buffer[s.len]
	s.buffer = s.buffer[:s.len]
	s.len--
	return ps
}

func (s *Stack) Reverse() string {
	if len(s.buffer) == 0 {
		return ""
	}
	ps := s.buffer[0]
	s.buffer = s.buffer[1:]
	s.len--
	return ps
}

func (s *Stack) Push(in ...string) bool {
	for i := range in {
		s.len++
		s.buffer = append(s.buffer, in[i])
	}
	s.len--
	return true
}

func (s *Stack) Len() int {
	return s.len
}

func Test_Stack(t *testing.T) {
	s := &Stack{}
	s.Push("a", "b", "c", "d", "e")
	for s.len >= 0 {
		t.Log(s.Reverse())
	}
}

type InputBuffer struct {
	Stack
}

func (i *InputBuffer) GetText() string {
	ret := ""
	for i.len >= 0 {
		ret += i.Reverse()
	}
	return ret
}

func (i *InputBuffer) RollBack() {
	i.Pop()
}
