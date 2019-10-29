package maxpoints

import (
	"testing"
)

func TestMaxPoints(t *testing.T) {
	arr := [3][2]int{{1, 2}, {4, 5}, {6, 7}}
	for i := range arr {
		t.Log(arr[i][0])
	}
	kmap := make(map[int]int)
	kmap[0] += 1
	kmap[0] += 1
	t.Log(kmap[0])
}
func TestCocows(t *testing.T) {
	res := CountCows(8)
	t.Log(res)
}

func CountCows(n uint) uint {
	// 为了降低时间复杂度 做了以下简洁运算
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 2
	}
	if n == 5 {
		return 4
	}
	if n == 6 {
		return 3
	}
	// 先声明6个数组
	res := [6]uint{0, 2, 0, 1, 0, 0} // 先赋值
	var i uint
	for i = 6; i < n; i++ {
		var newCows uint = 0
		if res[1] > 0 {
			newCows = res[1]
		}
		if res[3] > 0 {
			newCows += res[3]
		}
		var k uint = 5
		// 这里进行数据搬移操作
		for ; k > 0; k-- {
			res[k] = res[k-1]
		}
		res[0] = newCows
	}
	// 第六个不要了
	var sum uint = 0
	// 统计前五个数组的数值 得到牛的头数
	for j := 0; j < 5; j++ {
		sum += res[j]
	}
	return sum
}
