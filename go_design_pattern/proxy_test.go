package go_design_pattern_test

import "testing"

type UserIFace interface {
	Create() error
	Name() string
}

type User struct{}

func (u *User) Create() error {
	panic("implement me")
}

func (u *User) Name() string {
	panic("implement me")
}

type UserProxy struct {
}

func (u *UserProxy) Create() error {
	user := &User{}
	// some other methods
	return user.Create()
}

func (u *UserProxy) Name() string {
	user := &User{}
	// some other methods
	return user.Name()
}

func Test_proxy(t *testing.T) {
	var IUser UserIFace
	IUser = new(UserProxy)
	_ = IUser.Create()
}

// 动态代理 运行时创建代理类
// 不事先为每个原始类编写代理类，而是在运行的时候，动态地创建原始类对应的代理类，然后在系统中用代理类替换掉原始类。
