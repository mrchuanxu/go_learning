import (
	"testing"
	"encoding/json"
	"fmt"
	// "reflect"
	"strings"
)

func swapNums(arr []int){
	i,j := 0,len(arr)-1
	if j<=1{
		return
	}
	for i<j{
		if arr[i]>=0{
			i++
		}
		if arr[i]<0{
			if arr[j]>=0{
				arr[i],arr[j] = arr[j],arr[i]
				i++
				j--
			}else{
				j--
			}
		}
	}
}
// momoso { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}=>
// {‘A’: 1,‘B’: {‘A’: 2,‘B’: 3,},‘CC’: {‘D’: {‘E’: 4,‘F’: 5}}} 
// step 1 先将上面的编程结构体 取出string 作为key
// step 2 将string取split(".")
// step 3 将split后的字符作为map[string]的key值 存入

func flattenToNested(str string)string{
	res := make(map[string]interface{})
	f:= make(map[string]interface{})
	err := json.Unmarshal([]byte(str),&f)
	if err!=nil{
		return fmt.Sprintf("wrong %s",err)
	}
	for key,val := range f{
		s:=strings.Split(key,".")
		// 持续存入一个map 持续生成一个map
		// 遇到不存在key 即生成一个新的map存储数据
		resPart := makeMap(s,val.(float64))
		// 递推公式
		for ikey,ival :=  range resPart{
			if rv,ok:=res[ikey];ok{
				// 存在相等，继续判断 ival
				inMap(rv.(map[string]interface{}),ival.(map[string]interface{}))
			}else{
				res[ikey]=ival
			}
		}
		// 然后再循环放入res内
		
	}
	if v,err1:=json.Marshal(res);err1==nil{
		fmt.Println(string(v))
	}
	return ""
}
// 递归实现一个两个map的比较
func inMap(ival1 map[string]interface{},ival2 map[string]interface{}){
	if ival1==nil||ival2==nil{
		return
	}
	// 然后层层深入
	for irkey,irval := range ival1{
		for ivkey,iival:=range ival2{
			if ivkey==irkey{
				inMap(irval.(map[string]interface{}),iival.(map[string]interface{}))
				break
			}else{
				ival1[ivkey] = iival
				return
			}
		}
	}
}
//  递归实现一个 连续嵌套的map
func makeMap(val []string,inum float64)map[string]interface{}{
	rmap := make(map[string]interface{})
	if len(val)==1{
		rmap[val[0]] = inum
		return rmap
	}
	rmap[val[0]] = makeMap(val[1:],inum)
	return rmap
}
