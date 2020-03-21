package ch12_test

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

type astruct struct {
	Num1 int64
	Num2 int64
}

var x struct {
	a bool
	b uintptr
	c []int
}

func TestReflect(t *testing.T) {
	t.Log(unsafe.Sizeof(x.a))
	t.Log(unsafe.Sizeof(x.b))
	t.Log(unsafe.Sizeof(x.c))
	t.Log(unsafe.Sizeof(x))
	t.Log(unsafe.Offsetof(x.c))
	// m := reflect.TypeOf("hwllp")
	// t.Log(m)
	// n := reflect.ValueOf("j")
	// if n.Kind() == reflect.String {
	// 	t.Log("yes")
	// }
	// t.Log(n.Kind())
	// a := 2
	// c := reflect.ValueOf(&a)
	// d := c.Elem().Addr().Interface().(*int)
	// *d = 10
	// t.Log(a)

	// b := 3
	// e := reflect.ValueOf(&b).Elem()
	// e.Set(reflect.ValueOf(42))
	// t.Log(e.Type().Field(1))

	as := astruct{1, 2}
	bs := reflect.ValueOf(as)
	t.Log("bs type:", bs.Type())
	t.Log(subsets([]int{1, 2, 3}))
	t.Log(combine(4, 2))
	str := "ppp"
	t.Log(byte('A'), byte('Z'))
	t.Log(reflect.TypeOf(str[0]))
	s := []int{1, 2}
	s = s[0:0]
	t.Log(len(s))

}

func combine(n int, k int) [][]int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	ans := [][]int{}

	// 经典回溯算法
	// 从数组中寻找
	// 1，2 1，3 1，4
	var f func(ret *[][]int, arrs, prev []int, start, k int)
	f = func(ret *[][]int, arrs, prev []int, start, k int) {
		if len(prev) == k {
			*ret = append(*ret, append([]int{}, prev...))
			return
		}

		for i := start; i < len(arrs); i++ {
			f(ret, arrs, append(prev, arrs[i]), i+1, k)
		}

	}

	f(&ans, nums, []int{}, 0, k)
	return ans
}

func subsets(nums []int) [][]int {
	kLen := len(nums)
	if kLen == 0 {
		return nil
	}
	ans := [][]int{}

	var f func(ret *[][]int, arrs, prev []int, start, k int)
	f = func(ret *[][]int, arrs, prev []int, start, k int) {
		if len(prev) == k {
			*ret = append(*ret, append([]int{}, prev...))
			return
		}
		for i := start; i < len(arrs); i++ {
			f(ret, arrs, append(prev, arrs[i]), i+1, k)
		}
	}
	for i := 0; i < kLen; i++ {
		f(&ans, nums, []int{}, 0, i)
	}
	return ans
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	// Build map of fields keyed by effective name
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s:%v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s:%v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

func combinationSum(candidates []int, target int) [][]int {
	// 是一个排列组合问题
	// 回溯法解决
	if len(candidates) == 0 {
		return nil
	}

	ans := [][]int{}
	current := []int{}
	combination(ans, current, candidates, 0, target)
	return ans
}

func combination(ans [][]int, current []int, candidates []int, last_use, target int) {
	if target == 0 {
		ans = append(ans, current)
		return
	}
	for i := last_use; i < len(candidates) && candidates[i] <= target; i++ {
		current = append(current, candidates[i])
		combination(ans, current, candidates, i, target-candidates[i])
		current = current[:len(current)-1]
	}
}
