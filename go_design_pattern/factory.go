package go_design_pattern

// 将实现方式集中在一个方法 简单工厂方法
func SimpleFactory(a string) string {
	if a == "1" {
		return "1"
	}
	if a == "2" {
		return "2"
	}
	// ...
	return ""
}

// 简单工厂模式
var a1 = "1"
var a2 = "2"
var a3 = "3"

func SimpleFactoryM(a string) string {
	var m map[string]string
	m["1"] = a1
	m["2"] = a2
	m["3"] = a3

	if _, ok := m[a]; ok {
		return m[a]
	}
	return ""
}

type MainFactory interface {
	create() error
}

type SubFactory struct{}

func (s *SubFactory) create() error {
	return nil
}

type Sub1Factoty struct{}

func (s *Sub1Factoty) create() error {
	return nil
}

func CreateFactory() MainFactory {
	return &Sub1Factoty{}
}
