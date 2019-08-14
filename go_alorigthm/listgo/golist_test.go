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
	if li==nil{
		li = new(singleNode) // 当li为空，新建一个新的singleNode
		li.a = ch
		li.next = nil
		return li
	}
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

// 单链表反转 非递归
func reverseList(li *singleNode)(*singleNode,bool){
	if li == nil{
		return li,false
	}
	var pNew *singleNode
	var pTmp *singleNode
	for li!=nil{
		pTmp = li.next
		li.next = pNew
		pNew = li
		li = pTmp
	}
	li = pNew
	return li,true
}

// 单链表递归
func reverseListRecur(li *singleNode,revNode *singleNode)(*singleNode){
	if li == nil{
		return li
	}
	tmp := li.next
	li.next = revNode
	if tmp !=nil{
		return reverseListRecur(tmp,li)
	}else{
		return li
	}
}
// 链表环的检测 快慢指针 走链表
func circleList(li *singleNode)(bool){
	// if pFast == pSlow isExist Circle
	if li == nil{
		return false
	}

	var pFast *singleNode = li.next.next
	var pSlow *singleNode = li
	for pSlow != pFast{
		pFast = pFast.next.next
		pSlow = pSlow.next
		if pSlow == nil || pFast == nil{
			return false
		}
	}
	return true
}
// 两个有序链表合并 递归 贼快
func recurMerge(l1 *singleNode,l2 *singleNode)(*singleNode){
	if l1==nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var tmp *singleNode
	if l1.a < l2.a{
		tmp = l1
		tmp.next = recurMerge(l1.next,l2) // 递归寻值
	}else{
		tmp = l2
		tmp.next = recurMerge(l1,l2.next)
	}
	return tmp
}
// 删除倒数第n个节点 k-n+1
// func delbutN(li *singleNode,n int)(*singleNode,bool){

// }
// 求链表的中间结点 快慢指针 两倍 可以找到中间结点 不写了 翻篇
func TestSingleNode(t *testing.T){
	l1 := initList()
	l2 := initList()
	// selCh,_ := selectNode(2,li)
	arr := [10]rune{'a','c','g','h','i','j'}
	arr2 := [10]rune{'b','d','e','f','k','l'}
	for _,iarr := range(arr){
		l1 = insertRail(iarr,l1)
	}
	for _,iarr2 := range(arr2){
		l2 = insertRail(iarr2,l2)
	}
	selCh,_:= selectNode(4,l1)
	t.Logf("%c %c",selCh,l1.a)
	// isInsert:= insertList('a',li,-7)
	// for li.next!=nil{
	// 	t.Logf("%c",li.a)
	// 	li = li.next
	// }
	// l3 := recurMerge(l1,l2)

	// for l3!=nil{
	// 	t.Logf("%c",l3.a)
	// 	l3 = l3.next
	// }

}