// LRU缓存算淘汰策略

package listgo

import (
	"testing"
)

type singleNode struct{
	a rune // one char
	next *singleNode
}

func initList()(*singleNode){
	head := new(singleNode) // 动态申请内存空间
	head.next = nil
	return head
}
// 先建立链表 头插法 尾插法
func insertHead(ch rune,li *singleNode)(*singleNode){
	newHead := new(singleNode)
	newHead.next = li
	newHead.a = ch
	return newHead
}

func insertRail(ch rune,li *singleNode)(*singleNode){
	initLi := li
	nextHead := new(singleNode)
	for li.next != nil{
		li = li.next
	}
	li.next = nextHead
	nextHead.a = ch
	return initLi
}

// 一个示例 表示插入到第几个
func insertList(ch rune,li *singleNode,n int)bool{
	insertHead := new(singleNode)
	insertHead.a = ch
	i := 0
	for ;i<n&&li.next!=nil;i++{ // 考虑边界条件
		li = li.next // 去到插入的位置，然后插入
	}
	if i<n||n<0{ // n过大
		return false
	}
	if li.next ==nil{
		li.next = insertHead
		return true
	}
	insertHead.next = li.next
	li.next = insertHead
	return true
}

// 一个示例 表示删除到第几个
func deleteNode(li *singleNode,n int)(*singleNode,bool){
	initHead := li
	i:=0
	for ;i<n&&li.next!=nil;i++{
		li = li.next
	}
	if i<n||n<0{// n过大
		return li,false
	}
	recordNode := li.next
	li.next = recordNode.next
	recordNode = nil // 垃圾回收
	return initHead,true
}

// 表示链表查找第几个
func selectNode(n int,li *singleNode)(rune,bool){
	i:=0
	for ;i<n&&li.next!=nil;i++{
		li = li.next
	}
	if i<n||n<0||li==nil{
		return 'n',false
	}
	return li.a,true
}

// 表示修改第几个
func updateNode(n int,upCh rune,li *singleNode)(rune,bool){
	i:=0
	for ;i<n&&li.next!=nil;i++{
		li = li.next
	}
	if i<n||n<0||li==nil{
		return 'n',false
	}
	li.a = upCh
	return li.a,true
}

func TestSingleNode(t *testing.T){
	li := initList()
	arr := [10]rune{'a','b','c','d','e','f','g','h','i','j'}
	for _,iarr := range(arr){
		li = insertRail(iarr,li)
	}
	selCh,_:= selectNode(3,li)

	t.Logf("%c",selCh)

	isInsert:= insertList('a',li,-7)
	t.Log(isInsert)
	for li.next!=nil{
		t.Logf("%c",li.a)
		li = li.next
	}

}