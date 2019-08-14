package stackgo

import (
	"testing"
	"errors"
)

// 栈 一个操作有限的线性表 一个切片
type iStack []interface{}

// 长度
func (ista iStack) Len()int{
	return len(ista)
}

// 容量
func (ista iStack) Cap()int{
	return cap(ista)
}

// push
func (ista *iStack) push(value interface{})bool{
	*ista = append(*ista,value) // 解引用获取值 自动扩容 无需考虑大小
	return true
}

// pop
func (ista *iStack) pop()error{
	iStaValue :=*ista
	if len(iStaValue) == 0{
		return errors.New("out of range len is 0")
	}
	*ista = iStaValue[:len(iStaValue)-1] // pop就应该有出现值
	return nil
}

// top
func (ista iStack) top()(interface{},error){
	if len(ista) == 0{
		return nil,errors.New("No thing in Stack")
	}
	return ista[len(ista)-1],nil
}

// isEmpty
func (ista iStack) isEmpty()bool{
	return len(ista)==0
}

func TestStack(t *testing.T){

}


