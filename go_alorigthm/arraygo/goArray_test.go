// 数组为什么从0开始
// 为什么数组要从0开始编写，而不是从1开始呢？<br>
// 如何实现随机访问？
// 数组是一种线性数据结构，他用一组连续的内存空间，来存储一组具有相同类型的数据
// 一 线性表
// 二 连续的内存空间和相同类型的数据 随机访问，但删除插入 需要进行大量的数据搬移工作
// 根据下标随机访问的时间复杂度O(1)
// 

package arraygo

import (
	"testing"
)

func TestArray(t *testing.T){
	// i:=0
	// var arr [3]int
	// for ;i<=3;i++{ // panic 爆出out of range
	// 	arr[i] = 0;
	// 	t.Log("hello world\n")
	// }

	arr1 := [...]int{1,2,3,4}
	for iarr := range(arr1){
		t.Log(iarr)
	}

	var arr int = 1
	t.Log(arr)

	arr2 := [...]rune{'a','b','c','d'}
	for _,iarr2 := range(arr2){
		t.Logf("%c\n",iarr2)
	}
}