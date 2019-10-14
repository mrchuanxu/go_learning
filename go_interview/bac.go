package sol1

import (
	"testing"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
	"context"
)

func TestSol(t *testing.T){
	data := []int{1,2,3}
	// 传入的值就是值 不变
	for _,v := range data{
		v*=10
	}
	t.Log("data:",data)
}


// 数组传的是值 切片传的是引用！ 切片 capactiy与len 两个部分 传的是引用 capacity是2的指数增长 len就是值的多少
func change_slice(data [5]int){
	data[0] = 0
}

func TestSol1(t *testing.T){
	slice := [...]int{1,2,3,4,5}
	change_slice(slice)
	t.Log(slice)
}
// golang make与new的区别

type Student struct{
	name string
	age int
}

type Teacher interface{
	Name() string
	Action() string
}

func (s *Student) Name()string{
	fmt.Println("i am student trans")
	return "trans"
}

func (s *Student)Action()string{
	fmt.Println("i am a player")
	return "player"
}

func TestMake(t *testing.T){
	// 细究make与new区别
	aNew := Student{"momo",23}
	t.Log(reflect.TypeOf(aNew))
	t.Log(aNew.Name())
}


func TestRun(t *testing.T){
	runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i:=0;i<5;i++{
		go func(){
			fmt.Println("i:",i)
			wg.Done()
		}()
	}
	wg.Wait()
	// var list = make([]int,5)
	// t.Log(cap(list))
	// if AppendSlice(list,1){
	//     fmt.Println(list)
	// }
	// deferTest()
	ctx, cancel := context.WithCancel(context.Background())
    go my_task(ctx)
    time.Sleep(10 * time.Second)
    cancel()
    fmt.Println("cancel task")
    time.Sleep(2 * time.Second)
}
func my_task(ctx context.Context){
    ctx.Done() // 上下文 消灭子节点
}

func AppendSlice(list []int,i int)bool{
	if len(list)>10{
		fmt.Println("big")
		return true
	}else{
		list = append(list,i) // 申请的是5，相当于另外的引用 所以这个改变了原来的list，另起炉灶了
		return AppendSlice(list,i+1)
	}
}

func deferTest(){
	i:=1
	defer fmt.Println(i) // 从这里可以看出 i的值是已经赋值完了。只不过是延迟执行
	i++
	fmt.Println("iiiii",i)
	i++
	panic("exit")
	defer fmt.Println(i)
}


type data struct{
	name string
}

func TestSol2(t *testing.T){
	// m := map[string]data{"x":{"Tom"}}
	// m["x"].name = "Fuck"
	// 修改为以下
	m := make(map[string]*data)
	m["x"] = &data{"jerry"}
	m["x"].name  = "爱我中华"
	t.Log(m["x"].name)
}