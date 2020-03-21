package make_new_test

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 66, 7, 7, 88, 9}
	for i := range arr {
		t.Log("log", arr[i])
	}
	ForRange(arr)
}

// 递归实现循环
func ForRange(arr []int) {
	if len(arr) == 0 {
		return
	}
	fmt.Println("print", arr[0])
	ForRange(arr[1:])
}
