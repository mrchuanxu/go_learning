package binaryTreeSearch

import (
	"testing"
)

type searchNode struct{
	val rune
	left *searchNode
	right *searchNode
}

func (bTree *searchNode)findData(val rune)interface{}{
	if bTree == nil{
		return nil
	}
	for bTree != nil{
		if bTree.val > val{
			bTree = bTree.left
		}else if bTree.val < val{
			bTree = bTree.right
		}else{
			return bTree
		}
	}
	return nil
}

// 需要理解interface
func (bTree *searchNode)insertData(val rune)interface{}{
	if bTree ==nil{
		return nil
	}
	for bTree!=nil{
		if bTree.val>val{
			if bTree.left == nil{
				newNode := new(searchNode)
				newNode.val = val
				bTree.left = &newNode
			}
			bTree = bTree.left
		}else if bTree.val<val{
			if bTree.right == nil{
				newNode := new(searchNode)
				newNode.val = val
				bTree.right = &newNode
			}
			bTree = bTree.right
		}
	}
	return true
}