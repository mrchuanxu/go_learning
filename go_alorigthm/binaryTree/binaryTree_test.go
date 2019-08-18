package binaryTree

import (
	"testing"
	"fmt"
)

// 完全二叉树适合用数组存储
// 如何存储一棵二叉树 链式存储法和数组存储法

type treeNode struct{
	val interface{}
	left *treeNode
	right *treeNode
}

// 基于数组的顺序存储法 
// [][A][B][C][D][E][F][G][H][I][J][K]
//  0 1  2  3  4  5  6  7  8  9  10 11
//    i  2i 2i+1 ... 

// 遍历 中序遍历 前序遍历 后序遍历
// 递归实现 循环实现
func preOrder(root *treeNode){
	if root == nil{
		return
	}
	fmt.Println(root.val)
	preOrder(root.left)
	preOrder(root.right)
}

func inOrder(root *treeNode){
	if root == nil{
		return
	}
	inOrder(root.left)
	fmt.Println(root.val)
	inOrder(root.right)
}

func backOrder(root *treeNode){
	if root == nil{
		return
	}
	backOrder(root.left)
	backOrder(root.right)
	fmt.Println(root.val)
}

// 循环实现
func preOrder2(root *treeNode){
	// 其实很简单 为了避免递归过深的问题
	if root == nil{
		return
	}
	ista := new(iStack)
	ista.push(root)
	for ista.Len()!=0{
		isNode,err := ista.top().(*treeNode) // 进行类型断言
		if err{
		    fmt.Println(isNode.val)
		}else{
			panic("Such node type is wrong")
		}
		poperr:=ista.pop()
		if poperr!=nil{
			panic(poperr)
		}
		if isNode.right != nil{
			ista.push(isNode.right)
		}else if isNode.left != nil{
			ista.push(isNode.left)
		}
	}
}

// 循环实现 中序遍历
func inOrder2(root *treeNode){
	if root == nil{
		return
	}

	ista := new(iStack)
	ista.push(root)
	for ista.Len()!=0{
		isNode,err := ista.top().(*treeNode)
		if (isNode.left != nil)&&err{
			ista.push(isNode.left)
		}else{
			tmpNode,tmperr := ista.top().(*treeNode)
			fmt.Println(tmpNode.val)
			pterr := ista.pop()
			if pterr!=nil{
				panic(pterr)
			}
			if (tmpNode.right!=nil)&&tmperr{
				ista.push(tmpNode.right)
			}
		}
	}
	return
}

func TestOrder(t *testing.T){

}