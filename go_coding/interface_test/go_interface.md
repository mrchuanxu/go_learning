## 我想跟你聊聊interface{}
本来是写C++，写不下去了，做不了什么项目，搞个蛋！还是写go吧，go至少和现在的很多项目都相关。今天来讲讲go的interface。概念还是要说的，不然就会偏了，先来看看go文档对interface的解析 "An interface is a great and only way to achieve Polymorphism in Go" 嗯，是英文。作为一个程序员，英文是要懂的，不然跟不上潮流，你跟不上潮流就只能去学C++了😄。没错，interface就是设计来实现面向对象编程的，在面向对象中，C++就是继承，但是C++没有接口这个概念，用的是虚函数来实现接口的作用。虚函数还包括虚函数表，指针指向获取指定的虚函数，好了，不扯这个，扯一下interface
```go
type Shape interface{
    Area() float64
    Perimeter() float64
    Circle(width ...interface{}) float64
}
```
这就是go中的接口的声明方式，没有任何的实现迹象。emmm，那我们看看这个接口到底是个什么东西？

```go
var s Shape
fmt.Println(s)
fmt.Printf("type of s is %T\n",s)
```
竟然是nil，竟然是nil，出乎意料吧，来看看，interface其实有两个类型，一个是静态类型，也就是自己interface，一个类型指向动态值，也就是指针一样，就是下面这个图

sequenceDiagram
participant static_type
participant dynamic_value
static_type->dynamic_value:we build interface

为什么会是nil呢？因为这个interface都不知道自己应该指向什么类型，不知道继承什么，也就是没有赋值，fmt就是打印出这个动态值的，也就是nil了，但是实际上，这个interface是Shape。

来看看这个继承
```go
// 建一个Rect
type Rect struct{
	width float64
	height float64
}

func (r Rect)Area()float64{
	return r.width * r.height
}

func (r Rect)Perimeter()float64{
	return 2*(r.width+r.height)
}

func (r Rect)Circle(width ...float64)float64{
	if len(width)==0{
		return 0
	}
	isum := 0.0
	for _,iwidth := range(width){
		isum +=iwidth
	}

	return isum
}
```
这里的动态值就是Rect了，打印出来的T也就是Rect了。动态类型是指，每个实例实现了这个interface，例如我加一个Quator的struct，然后按照上面来实现interface所有的方法，然后将Quator赋值给s，s打印出来就是Quator，这就是我们说的动态类型。在C++是这样，在这里也是这样，但是一定要全部实现interface的方法，不然，编译会错误，当然，struct可以实现interface以外的方法。<br>

那么空的interface呢？空的？
### interface{}
这个有点像，java中的Object，所有方法都继承自Object，所以这个也是一样的，他是空的，所以，所有的类型都继承于这个空的interface{}，无招胜有招嘛！来看看一个例子，fmt的Println是怎么打印不同类型的data的？就是用的empty interface！

```go
type MyString string

func explain(i interface{}){
	
	// 要是这里使用i 就需要断言判断了
	switch val :=i.(type){
	case string:
		fmt.Println(val)
	case int:
		val = val+1
		fmt.Println(val)
	default:
		fmt.Printf("%v,%T",i,i)
	}
}
```
好了，是时候放大招了
### Multiple interfaces
一个类型是可以继承多个interface的，

```go
type Shape interface{
    Area()float64
}
type Object interface{
    Volume() float64
}

type Cube struct{
    side float64
}
...// 然后就是实现这两个接口了 这就不写了吧

func main(){
    c := Cube{3}
    var s Shape = c
    var o Object = c
    fmt.Println("Volume of o of interface type Shape is", o.Volume())
    fmt.Println("Area of s of interface type Object is", s.Area())
}
```
上面的写法是ok的，那么调过来呢？例如这样
```go
fmt.Println("area of s of interface type Shape is", s.Volume())
fmt.Println("volume of o of interface type Object is", o.Area())
```
这就错误了，没定义过呀，因为s的静态类型是Shape，o的静态类型是Object。但是，要是想让这个可以工作呢？
### Type assertion
动态值可以用断言i.(type)获得指向的类型以及继承的类型，go会检查，这个动态类型是不是这个类型。嗯。很直接(Go will check if dynamic type of i is identical to Type)

```go
var s Shape = Cube{3}
c := s.(Cube) // 这里我判断是不是这个类型，返回的是什么？
```
断言提取了指向Cube的类型，所以c就名正言顺滴使用Cube的方法。其实还有另外一种语法保证了提取的安全
```go
value,ok := i.(Type)
```
ok表示这个Type是不是继承了i类型，然后如果是，那么ok就是true，反之则反。

### Embedding interfaces
可以合并
```go
type merge interface{
    Shape
    Object
}
...// 然后直接继承使用
```
### 指针还是值？
```go
// 建一个Rect
type Rect struct{
	width float64
	height float64
}

func (r *Rect)Area()float64{
	return r.width * r.height
}

func (r *Rect)Perimeter()float64{
	return 2*(r.width+r.height)
}

func (r Rect)Circle(width ...float64)float64{
	if len(width)==0{
		return 0
	}
	isum := 0.0
	for _,iwidth := range(width){
		isum +=iwidth
	}

	return isum
}

func main(){
	r := Rect{5.0,4.0}
	var s Shape = r // ❌wrong
	// program.go:27: cannot use Rect literal (type Rect) as type Shape in assignment: Rect does not implement Shape (Area method has pointer receiver)
}
```
what the hell???黑人问号脸，我这只是变了指针，为什么会报错？Rect确实是继承了Shape的所有方法了呀。在这个接口中，一旦一个方法有了指针，这个接口就存在一个动态类型指针，而不是动态类型的值，因此，赋值应该是直接赋予指针值，而不是值。所以这样就对了`var s Shape = &r`。<br>
#### 总结？！
好了。讲完了。啊？go的多态就讲完了？哦，没有没有，还有几个概念，why interface
* writing generic algorithm（泛型编程）
* hiding implementation detail
* providing interception points

这几个应该都很容易懂，C++实现范型编程用的是模版，参照STL标准库可以看出，实现还是蛮简单的，但是go实现泛型编程就是使用的interface，其实我们可以看一下interface底层的结构，C++

```go
type eface struct{
	_type *_type
	data unsafe.Pointer
}
   type _type struct {
        size       uintptr // type size
        ptrdata    uintptr // size of memory prefix holding all pointers
        hash       uint32  // hash of type; avoids computation in hash tables
        tflag      tflag   // extra type information flags
        align      uint8   // alignment of variable with this type
        fieldalign uint8   // alignment of struct field with this type
        kind       uint8   // enumeration for C
        alg        *typeAlg  // algorithm table
        gcdata    *byte    // garbage collection data
        str       nameOff  // string form
        ptrToThis typeOff  // type for pointer to this type, may be zero
    }
```
根据interface是否包含method，底层实现上用struct表示，iface和eface，eface表示不含method的interface的结构，或者叫empty interface。对于Golang中的大部分数据类型都是可以抽象出来的_type结构，同时针对不同的类型还是会有一些其他信息
那么iface就是这样的，不是空的interface<br>
```go
    type iface struct {
        tab  *itab
        data unsafe.Pointer
    }
    
    // layout of Itab known to compilers
    // allocated in non-garbage-collected memory
    // Needs to be in sync with
    // ../cmd/compile/internal/gc/reflect.go:/^func.dumptypestructs.
    type itab struct {
        inter  *interfacetype
        _type  *_type
        link   *itab
        bad    int32
        inhash int32      // has this itab been added to hash?
        fun    [1]uintptr // variable sized
    }
```
编译器检测，struct是不是满足interface的类型要求。其实，通过了解内存布局哦，可以说，能够进一步分析类型断言等情况的效率问题。当判定一种类型是否满足某个接口时，golang使用类型的方法集和接口所需要的方法集进行匹配，如果类型的方法集完全包含接口的方法集，则可认为该类型满足该接口。

### 不讲了，剩下自己开发继续使用吧！
* [参考知乎优秀回答](https://www.zhihu.com/search?type=content&q=golang%20interface%7B%7D)
