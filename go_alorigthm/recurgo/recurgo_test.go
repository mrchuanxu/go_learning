package recurgo

import (
	"testing"
)
// 1. 一个问题可以分成几个子问题
// 2. 分解之后的子问题，除了数据规模不同，求解思路完全相同
// 3. 存在递归终止条件
// 关键写递推公式 找到终止条件

var depth int = 0
func recu(n int)int{
	depth++
	if depth>3{
		panic("too depth")
	}
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	return recu(n-1)+recu(n-2)
}
// 写递归代码的关键就是找到如何将大问题分解成小问题的规律，并且基于此写出递推公式
// 然后再推敲终止条件 最后将递推公式和终止条件翻译成代码
// 函数调用会使用栈来保存临时变量
func TestRecu(t *testing.T){
	inum := recu(10)
	t.Log(inum)
}