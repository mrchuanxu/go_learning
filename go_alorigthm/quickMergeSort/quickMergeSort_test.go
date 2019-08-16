package quickMergeSort

import (
	"testing"
	"fmt"
	"math/rand"
)

// 先写递推公式
// 再写具体实现
// 
func mergeSort(arr []rune,start int,end int){
	if start >= end{
		return
	}
	mid := (start+end)/2
	mergeSort(arr,start,mid)
	mergeSort(arr,mid+1,end)
	goMerge(arr,start,mid,end)
}

func goMerge(arr []rune,start int,mid int,end int){
	i,j := start,mid+1
	tmp := []rune{}
	for i<=mid&&j<=end{
		if arr[i]<arr[j]{
			tmp = append(tmp,arr[i])
			i++
		}else{
			tmp = append(tmp,arr[j])
			j++
		}
	}
	// 左边没走完
	for i<=mid{
		tmp = append(tmp,arr[i])
		i++
	}
	// 右边没走完
	for j<=end{
		tmp = append(tmp,arr[j])
		j++
	}
	// 这里直接计算tmp的长度
	for l:=0;l<len(tmp);l++{
		arr[start+l] = tmp[l]
	}
	tmp = nil
}

// 快排 中心思想 获得分区点
func quickSort(arr []rune,start int,end int){
	if start >= end{
		return
	}
	q := partition(arr,start,end)
	quickSort(arr,start,q-1)
	quickSort(arr,q+1,end)
}
// 	优化快速排序 三点取值
// func threePickOne(arr[] rune,start int,end int)int{

// }
// 先排
func partition(arr []rune,start int,end int)int{
	// 分区 暂时取最后一个元素
	// 随机获取一个分区点 这样避免O(n^2)的退化
	valPivot := rand.Intn(end-start)+start
	pivot := arr[valPivot]
	// 交换
	i := start
	for j:=start;j<end;j++{
		if arr[j] < pivot{
			arr[i],arr[j] = arr[j],arr[i] // 类似选择排序 将已处理区间放在前面 未处理区间放在后面
			i++
		}
	}
	arr[i],arr[valPivot] = arr[valPivot],arr[i]
	return i
}

// 在O(n)的时间复杂度内求无序数组中的第k大元素
// 就是修改一下quickSort的函数即可
func theKNum(arr []rune,start int,end int,k int){
	if start>=end{
		return
	}
	p := partition(arr,start,end)
	if p+1 == k{
		fmt.Printf("%c",arr[p])
		return
	}else if k>p+1{
		theKNum(arr,p+1,end,k)
	}else{
		theKNum(arr,start,p-1,k)
	}
}



func TestMerge(t *testing.T){
	arr := []rune{'a','c','g','d','e','f','k'}
	quickSort(arr,0,len(arr)-1)
	for _,iarr := range(arr){
		t.Logf("%c",iarr)
	}
}