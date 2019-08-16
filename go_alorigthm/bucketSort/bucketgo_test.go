package bucketSort

import (
	"testing"
	"errors"
)
// 桶排序对要排序数据分要求是非常苛刻
// 计数排序 个人实现
// 传参用指针操作
func countSort(arr []int,max int)error{
	// 首先对数组求和 先放桶里面 指定切片的长度和容量
	if len(arr) <= 1{
		return errors.New("Nothing to sort")
	}
	toCount := make([]int,max+1)
	for _,iarr := range(arr){
		// 将数据放入桶内
		toCount[iarr]++
	}
	// 对toCount顺序求和
	for idx,tarr := range(toCount){
		if idx != 0{
			toCount[idx] = toCount[idx-1]+tarr
		}
	}
	toSort := make([]int,len(arr))
	// 将结果依次放入
	for _,itarr := range(arr){
		toCount[itarr]--
		toSort[toCount[itarr]] = itarr
	}
	// 将结果拷贝给arr
	for idx,itsArr := range(toSort){
		arr[idx] = itsArr
	}
	return nil
}

func TestBuckSort(t *testing.T){
	arr := []int{2,5,3,0,2,3,0,3}
	err := countSort(arr,5)
	t.Log(err)
	for _,iarr := range(arr){
		t.Log(iarr)
	}
}