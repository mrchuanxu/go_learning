package sortgo

import (
	"testing"
	"errors"
)

// swap 没意思 没有引用传递 就像没法搞事情
// func Swap(a int,b int)(int,int){
// 	a,b = b,a
// 	return a,b
// }
// 冒泡排序
func bubbleSort(arr []rune)error{
	// i,j 对比
	arrLen := len(arr)
	if arrLen == 0{
		return errors.New("Nothing to Sort,len is 0")
	}
	flag := true
	for i:=0;i<arrLen;i++{
		for j:=0;j<arrLen-i-1;j++{ // 最上面一层就不需要排了
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				flag = true
			}
			if !flag{
				break
			}
		}
	}
	return nil
}

func InsertSort(arr []rune)error{
	// 已排序区间与未排序区间
	arrLen := len(arr)
	if arrLen <= 1{
		return errors.New("Nothing to sort")
	}
	for i:=1;i<arrLen;i++{
		val := arr[i]
		j := i-1
		for ;j>=0;j--{
			if arr[j]>val{
				arr[j+1] = arr[j]
			}else{
				break
			}
		}
		arr[j+1] = val
	}
	return nil
}

// 选择排序
func selectSort(arr []rune)error{
	arrLen := len(arr)
	if arrLen <=1{
		return errors.New("Nothing to sort")
	}
	for i:=0;i<arrLen;i++{
		minVal := arr[i]
		k:=i // 记录交换
		for j:=i+1;j<arrLen;j++{
			if minVal>arr[j]{
				minVal = arr[j]
				k = j
			}
		}
		// 循环结束 找到最小的位置 交换
		arr[i],arr[k] = arr[k],arr[i]
	}
	return nil
}


func TestSwap(t *testing.T){
	arr := []rune{'1','4','6','4','3'}
	err:= selectSort(arr)
	t.Log(err)
	for _,iarr := range(arr){
		t.Logf("%c",iarr)
	}
}

