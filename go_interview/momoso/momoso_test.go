package momoso_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func arrSort(arr []int) {
	i, j := 0, len(arr)-1
	if j <= 1 {
		return
	}
	for i < j {
		if arr[i] >= 0 {
			i++
		}
		if arr[i] < 0 {
			if arr[j] >= 0 {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			} else {
				j--
			}
		}
	}
}

// "{ ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}"
// {'A':1,'B':{'A':2,'B':3},"CC":{"D":{"E":4,"F":5}}}

var strMap = map[string]string{"1": "bar", "2": "foo.bar", "3": "foo.foo", "4": "baz.cloudmall.com", "5": "baz.cloudmall.ai"}

func UnFlatten(sm map[string]int) (ret map[string]interface{}) {
	// 一个遍历
	var f func(s []string, index, val int, tmp map[string]interface{})
	f = func(s []string, index, val int, tmp map[string]interface{}) {
		if len(s)-1 == index {
			tmp[s[index]] = val
			return
		}
		if _, exist := tmp[s[index]]; !exist {
			mtmp := map[string]interface{}{}
			tmp[s[index]] = mtmp
			tmp = mtmp
		} else {
			tmp = tmp[s[index]].(map[string]interface{})
		}
		f(s, index+1, val, tmp)
	}
	ret = make(map[string]interface{})
	retTmp := ret
	for k, v := range sm {
		// 先切分
		retTmp = ret
		cm := strings.Split(k, ".")
		if len(cm) == 1 {
			if _, exist := retTmp[cm[0]]; !exist {
				retTmp[cm[0]] = v
			}
			continue
		}
		f(cm, 0, v, retTmp)
	}
	return
}

func CloudmallInterview1(numbers []int) []int {
	i, j := 0, len(numbers)-1
	if j <= 1 {
		return numbers
	}
	for idx := 0; idx < len(numbers); idx++ {
		if numbers[idx] >= 0 {
			numbers[idx], numbers[i] = numbers[i], numbers[idx]
			i++
		}
	}
	return numbers
}

func TestCloudmallInterview1(t *testing.T) {
	nums := []int{6, 4, -3, 0, 5, -2, -1, 0, 1, -9}
	nums = CloudmallInterview1(nums)
	t.Log(nums)
}

func CloudmallInterview2(revList map[string]string) (jsonStr string, err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Fatalf("CloudmallInterview2 has err [%v] with revList[%+v]", p, revList)
		}
	}()
	// 一个遍历
	if revList == nil {

	}
	var f func(s []string, index int, val string, tmp map[string]interface{})
	f = func(s []string, index int, val string, tmp map[string]interface{}) {
		if len(s)-1 == index { // 到达最后一个点
			tmp[s[index]] = val
			return
		}
		if _, exist := tmp[s[index]]; !exist {
			mtmp := map[string]interface{}{}
			tmp[s[index]] = mtmp
			tmp = mtmp
		} else {
			tmp = tmp[s[index]].(map[string]interface{})
		}
		f(s, index+1, val, tmp)
	}
	ret := make(map[string]interface{})
	retTmp := ret
	for k, v := range revList {
		// 先切分
		retTmp = ret
		cm := strings.Split(v, ".")
		if len(cm) == 1 {
			if _, exist := retTmp[cm[0]]; !exist {
				retTmp[cm[0]] = k
			}
			continue
		}
		f(cm, 0, k, retTmp)
	}
	strByte, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}
	jsonStr = string(strByte)
	return
}

func TestCloudmallInterview2(t *testing.T) {
	str, err := CloudmallInterview2(strMap)
	t.Log(str, err)
}
