package hashTable

// 散列函数
import (
	"testing"
	"errors"
)
// 哈希解决
func hash(key rune)int{
	hashValue := int(key)%3
	return hashValue
}
// LRU缓存淘汰算法

type LRU interface{
	insertCache(val interface{})(*interface{},error)
	selAndDelCache(val int,status int)(*interface{},error)
}

const (
	MAX_LEN = 100
	MAX_DLEN = 1000
	HASH_LEN = 10
)
// 构建每个结点
type singNode struct{
	val interface{} // 这个接口类型真的什么都能存 底层是void*
	Count int // 拉链长度
	dCount int // 双向链表长度
	prev *singNode
	next *singNode
	hnext *singNode
}

// 构建一个散列表 结构
type hashTab struct{
	hashArr [HASH_LEN]*singNode // 声明
	// 坐标作为击中对象，存储值求余获取
}

// 新建一个节点
func newSingNode(val interface{},len int,dLen int)*singNode{
	return &singNode{val:val,Count:len,dCount:dLen,prev:nil,next:nil,hnext:nil}
}
// 哈希函数 string
func hashCalcu(val int)int{
	//求余较好
	return val%10
}

// 查一个hnext链表
func selectHnext(head *singNode,val interface{})(**singNode,error){
	for head != nil{
		if head.val == val{
			return &head,nil
		}
		head = head.hnext
	}
	return nil,errors.New("No this node")
}

// 双向链表 断链而后将节点放到最后 1 代表查找 0 代表删除
func selectDnext(head *singNode,status int)(*singNode,error){
	if head.next == nil||head == nil||head.prev == nil{
		return head,errors.New("No next node")
	}
	preHead := head.prev
	nxtHead := head.next
	preHead.next = nxtHead
	nxtHead.prev = preHead
	iDlen := head.dCount
	for nxtHead.next!=nil{ // 找到最后的节点
		iDlen++
		nxtHead.dCount--
		nxtHead = nxtHead.next
	}
	if status == 1{ // 查找
		nxtHead.next = head // 将当前节点放到最后
		head.prev = nxtHead
		head.next = nil
		head.dCount += iDlen
	}else{ // 删除
		head = nil
	}
	return head,nil
}
// 查找和删除一个数据 int
func (hashtab *hashTab)selAndDelCache(val int,status int)(*singNode,error){
	// 先用哈希计算这个值的位置
	//TODO: 暂时求int 其余哈希以后加
	pfHead := &(hashtab.hashArr[hashCalcu(val)])
	// 查找这个链表 hnext 查找
	phnext,err := selectHnext((*pfHead),val)
	if err != nil{
		return nil,errors.New("No this node")
	}
	// 然后把找到的点 双向链表放在最后
	return selectDnext((*phnext),status)
}

// 插入一个数据 int
func (hashtab *hashTab) insertCache(val int)(*singNode,error){
	// 首先判断是否是第一个元素
	key := hashCalcu(val)
	pfHead := &(hashtab.hashArr[key])
	// 看看是不是hnext第一个节点
	// 看看是不是双向链表的第一个节点
	// 双向链表的计数
	// pfHead = *pfHead
	currentNode := pfHead
	var newNode *singNode
	dCount := 0
	// 1 随机寻找一个点 找前后
	if (*pfHead) == nil{
		flag := false
		// 遍历10个数组指针 发现是否hnext
		for i:=0;i<HASH_LEN;i++{
			if hashtab.hashArr[i] != nil{
				flag = true
				currentNode = &(hashtab.hashArr[i])
				dCount = (*currentNode).dCount
				for (*currentNode).next != nil{
					dCount++
					(*currentNode) = (*currentNode).next
				}
				break
			}
		}
		if flag{
			newNode = newSingNode(val,0,dCount) // TODO: 计算双向链表的节点
			(*currentNode).next = newNode
			newNode.prev = (*currentNode)
			(*pfHead) = newNode
		}else{ // 最新的节点
			newNode = newSingNode(val,0,dCount)
			(*pfHead) = newNode
		}  
	}else{
		// 遍历hnext
		for (*pfHead).hnext != nil{
			(*pfHead) = (*pfHead).hnext
		}
		currentNode:=(*pfHead)
		for currentNode.next != nil{
			currentNode = currentNode.next
		}
		newNode = newSingNode(val,(*pfHead).Count+1,currentNode.dCount+1)
		(*pfHead).hnext = newNode
		currentNode.next = newNode
		newNode.prev = currentNode
	}
	return newNode,nil
}

func TestLRU(t *testing.T){
	hasArr := new(hashTab)
	t.Log(hasArr.hashArr)
	for i:=0;i<1000;i++{
		currentNode,err:=hasArr.insertCache(i)
		if err == nil{
			t.Log(currentNode.val)
		}
	}
	isNode,err := hasArr.selAndDelCache(776,1)
	t.Log(isNode.val,err)
}

