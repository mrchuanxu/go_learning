package binaryTree

import (
	"fmt"
	"testing"
)

// 完全二叉树适合用数组存储
// 如何存储一棵二叉树 链式存储法和数组存储法

type treeNode struct {
	val   interface{}
	left  *treeNode
	right *treeNode
}

// 基于数组的顺序存储法
// [][A][B][C][D][E][F][G][H][I][J][K]
//  0 1  2  3  4  5  6  7  8  9  10 11
//    i  2i 2i+1 ...

// 遍历 中序遍历 前序遍历 后序遍历
// 递归实现 循环实现
func preOrder(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preOrder(root.left)
	preOrder(root.right)
}

func inOrder(root *treeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.val)
	inOrder(root.right)
}

func backOrder(root *treeNode) {
	if root == nil {
		return
	}
	backOrder(root.left)
	backOrder(root.right)
	fmt.Println(root.val)
}

// 循环实现
func preOrder2(root *treeNode) {
	// 其实很简单 为了避免递归过深的问题
	if root == nil {
		return
	}
	ista := new(iStack)
	ista.push(root)
	for ista.Len() != 0 {
		isNode, err := ista.top().(*treeNode) // 进行类型断言
		if err {
			fmt.Println(isNode.val)
		} else {
			panic("Such node type is wrong")
		}
		poperr := ista.pop()
		if poperr != nil {
			panic(poperr)
		}
		if isNode.right != nil {
			ista.push(isNode.right)
		} else if isNode.left != nil {
			ista.push(isNode.left)
		}
	}
}

// 循环实现 中序遍历
func inOrder2(root *treeNode) {
	if root == nil {
		return
	}

	ista := new(iStack)
	ista.push(root)
	for ista.Len() != 0 {
		isNode, err := ista.top().(*treeNode)
		if (isNode.left != nil) && err {
			ista.push(isNode.left)
		} else {
			tmpNode, tmperr := ista.top().(*treeNode)
			fmt.Println(tmpNode.val)
			pterr := ista.pop()
			if pterr != nil {
				panic(pterr)
			}
			if (tmpNode.right != nil) && tmperr {
				ista.push(tmpNode.right)
			}
		}
	}
	return
}

func TestOrder(t *testing.T) {

}

type MedianFinder struct {
	BigHeap []int
	LowHeap []int
}

// 主要是建堆 添加数字 分为大顶堆和小顶堆
// 需要不断重新建堆 最后大顶堆和小顶堆堆顶数字 取中位数
// 数据计数器为偶数时 插入最小堆 为奇数时 插入最大堆，
// 最大堆 当插入元素 比堆顶的要小 则取出堆顶元素 插入最小堆
// 最小堆 当插入元素 比堆顶元素大 则取出堆顶元素 插入最大堆
// 持续堆化

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{
		BigHeap: make([]int, 1),
		LowHeap: make([]int, 1),
	}
	mf.BigHeap[0] = int(^uint(0) >> 1)
	mf.LowHeap[0] = int(^uint(0) >> 1)
	return mf
}

func (this *MedianFinder) heapify() {
	// 取最后的数据来进行堆化
	// 最大堆堆化
	BLen := len(this.BigHeap) - 1
	for (BLen/2) > 0 && (this.BigHeap[BLen] > this.BigHeap[BLen/2]) {
		this.BigHeap[BLen], this.BigHeap[BLen/2] =
			this.BigHeap[BLen/2], this.BigHeap[BLen]
		BLen = BLen / 2
	}

	// 最小堆堆化
	LLen := len(this.LowHeap) - 1
	for (LLen/2) > 0 && (this.LowHeap[LLen] < this.LowHeap[LLen/2]) {
		this.LowHeap[LLen], this.LowHeap[LLen/2] =
			this.LowHeap[LLen/2], this.LowHeap[LLen]
		LLen = LLen / 2
	}
}

func (this *MedianFinder) AddNum(num int) {
	BLen, LLen := len(this.BigHeap)-1, len(this.LowHeap)-1
	if (BLen+LLen)&1 == 0 { // 偶数
		if len(this.BigHeap)-1 > 0 && this.BigHeap[1] > num {
			this.BigHeap = append(this.BigHeap, num)
			this.heapify()
			num = this.BigHeap[1]
			this.BigHeap = this.BigHeap[1:]
		}
		this.LowHeap = append(this.LowHeap, num)
		this.heapify()
		return
	}
	if len(this.LowHeap)-1 > 0 && this.LowHeap[1] < num {
		this.LowHeap = append(this.LowHeap, num)
		this.heapify()
		num = this.LowHeap[1]
		this.LowHeap = this.LowHeap[1:]
	}
	this.BigHeap = append(this.BigHeap, num)
	this.heapify()
	return
}

func (this *MedianFinder) FindMedian() float64 {
	SLen := (len(this.BigHeap) + len(this.LowHeap)) - 2
	if SLen == 0 {
		return 0
	}
	if SLen&1 == 1 {
		return float64(this.LowHeap[1])
	}
	return float64(this.BigHeap[1]+this.LowHeap[1]) / 2.0

}

func TestAddNum(t *testing.T) {
	mf := Constructor()
	mf.AddNum(-1)
	t.Log(mf.FindMedian())
}
