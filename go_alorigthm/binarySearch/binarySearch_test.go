package binarySearch

import (
	"testing"
)

func binarySearch(arr []int,left int,right int,val int)int{
	if left >= right{
		panic("not availiable") 
	}
	for left <= right{
		mid := left + ((right-left)>>1)
		if arr[mid] == val{
			return mid
		}else if arr[mid]<val{
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return -1
}

// recur
func binarySearchRecu(arr []int,left int,right int,val int)int{
	if left > right{
		return -1
	}
	mid := left + ((right-left)>>1)
	if arr[mid] == val{
		return mid
	}else if arr[mid]<val{
		return binarySearchRecu(arr,mid+1,right,val)
	}else{
		return binarySearchRecu(arr,left,mid-1,val)
	}
	return -1
}
// 求一个数的平方根

// 查找第一个值等于给定值的元素 只需要 加判断是否左边还有同样的值
func findFirstVal(arr []int,val int)int{
	low := 0
	high := len(arr)
	for low<=high{
		mid := low+((high-low)>>1)
		if arr[mid]>val{
			high = mid-1
		}else if arr[mid]<val{
			low = mid+1
		}else if (arr[mid-1]!=val)||(mid==0){
			return mid
		}else{
			high = mid-1
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素 只需要 加判断是否左边还有同样的值
func findLastVal(arr []int,val int)int{
	low := 0
	high := len(arr)
	arrLen := high
	for low<=high{
		mid := low+((high-low)>>1)
		if arr[mid]>val{
			high = mid-1
		}else if arr[mid]<val{
			low = mid+1
		}else if (arr[mid+1]!=val)||(mid==(arrLen-1)){
			return mid
		}else{
			high = mid-1
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T){
	
}