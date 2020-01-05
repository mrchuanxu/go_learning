package ch1_test

import (
	"testing"
)

// A.如何在给定的4300000000个32位顺序文件找出一个至少出现两次的整数
// 在40亿个不顺序排列整数中找一个不存在的整数 32位 一共4294967294个数
// 推荐二分搜索 惊人的搜索速度
// 现在是已经有4300000000 找出现两次的 那么就是5032706个至少出现两次的整数
// 顺序文件 二分查找 但是必须落在有至少出现两次的整数中
// B. 向量旋转算法 实现向量旋转 参考reverse
// C. 按键编码 9位按键 电话号码 匹配名字 返回名字参数
// 其实主要是要做标识，然后也是二分法 查找
func TestSearchInt(t *testing.T) {
	str := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	t.Log(str)
	reverse(str, 0, 2)
	reverse(str, 3, 7)
	reverse(str, 0, 7)
	t.Log(str)
}

func reverse(str []rune, i, n int) {
	for i <= n {
		str[i], str[n] = str[n], str[i]
		i++
		n--
	}
	return
}
