package ch6_test

import "testing"

import "bytes"

import "fmt"

// 一个bit数组通常会用一个无符号数或者称之为"字"的slice来表示
// 每一个元素的每一个位都表示集合里的一个值。当集合的第i位被设置时，我们才说这个集合包含元素i
//

var bitSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	// 当x为bitSize word 为1 若，len为1 则表示仅有一个位，2则有两个位 s.words的位置与bit左移求&
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	// 大于 等于
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// s.words {0}
	// bit  0
	s.words[word] |= 1 << bit
	// s.words[1] = 1
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/bitSize, uint(x%bitSize)
	s.words[word] &= uint(bit)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() (t *IntSet) {
	t = &IntSet{
		words: make([]uint, 10),
	}
	for i := range s.words {
		t.words[i] = s.words[i]
	}
	return
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) AddAll(x ...int) {
	for i := range x {
		word, bit := x[i]/bitSize, uint(x[i]%bitSize)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}
func TestBit(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if x.Has(9) {
		fmt.Println("has")
	}
	x.Remove(9)
	if !x.Has(9) {
		fmt.Println("not has")
	}
	y := x.Copy()
	if !y.Has(144) {
		fmt.Println("y not has")
	}
	fmt.Println(&x)
	fmt.Println(x)

	// var _ = IntSet{}.String() 不能在一个不能寻址的值调用string方法 若是对象就可以
	strRet := generateParenthesis(3)
	t.Log(strRet)

}

var strRet = []string{}

func generateParenthesis(n int) []string {
	generate("", 0, 0, n)
	return strRet
}

func generate(cur string, start, end, max int) {
	if len(cur) == 2*max {
		strRet = append(strRet, cur)
		return
	}

	if start < max {
		generate(cur+"(", start+1, end, max)
	}
	if end < start {
		generate(cur+")", start, end+1, max)
	}
}
