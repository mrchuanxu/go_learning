package ddd_test

type User struct {
	name      string
	mobile    string
	GetName   func() string
	GetMobile func() string
	U         IUser
}

type IUser interface {
	GetUser() string
}
