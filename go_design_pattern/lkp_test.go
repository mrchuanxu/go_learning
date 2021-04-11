package go_design_pattern_test

type Flyable interface {
	Fly(string)string
}

type Tweetable interface{
	Tweet(string)string
}

type LayEgg interface {
	Lay(string)string
}

type Flyablitity struct{}

func (f Flyablitity) Fly(s string) string {
	panic("implement me")
}

type Tweetablitity struct{}

func (t Tweetablitity) Tweet(s string) string {
	panic("implement me")
}

type Layablitity struct{}

func (l Layablitity) Lay(s string) string {
	panic("implement me")
}
// 组合实现 委托其他类实现
type Ostrich struct {
	Layablitity
	Tweetablitity
}

type Address struct{
	Province string
	City string
	Country string
}

type Person struct{ // 实体
	Name string
	Age int64
	Gender string
	address Address // 值对象
}


