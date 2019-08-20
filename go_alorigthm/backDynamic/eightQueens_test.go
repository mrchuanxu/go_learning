package backDynamic

import (
	"testing"
	"fmt"
)

// 8皇后问题 看懂回溯算法

var result [8]int

func cal8Queens(row int){
	if row == 8{ // 放完了 不放了
		print8Queens(result)
		return
	}
	for col:=0;col<8;col++{
		if isOK(row,col){
			result[row] = col
			cal8Queens(row+1) //  回溯
		}
	}
	return
}

// isOK? 放皇后！
func isOK(row,col int)bool{
	leftup := col-1
	rightup := col+1
	for i:=row-1;i>=0;i--{
		if result[i] == col{
			return false
		}
		// 考察对角线
		if leftup>=0{
			if result[i] == leftup{
				return false
			}
		}
		if rightup <8{
			if result[i] == rightup{
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

// print8Queens

func print8Queens(result [8]int){
	for row:=0;row<8;row++{
		for col := 0;col<8;col++{
			if result[row] == col{
				fmt.Print("Q")
			}else{
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}


// 0-1 背包问题
const (
	MaxUint =^uint(0)
)
var maxWeight int = -int(MaxUint >> 1)-1

func packMaxWeight(i,cw,w,n int,items []int){
	if i == n || cw==w{
		if cw > maxWeight{
			maxWeight = cw
		}
		return
	}
	packMaxWeight(i+1,cw,w,n,items) // 递归 不选择这个物品
	if (cw+items[i])<=w{
		packMaxWeight(i+1,cw+items[i],w,n,items) // 选择这个物品
	}
}


func TestQueens(t *testing.T){
	cal8Queens(0)
}
