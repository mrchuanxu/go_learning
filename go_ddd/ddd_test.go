package ddd_test

import (
	"testing"
)

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

func Test_Slice(t *testing.T) {
	sarr := []string{"1", "3", "4", "5"}
	for len(sarr) > 1 {
		tmpArr := sarr[:1]
		t.Log(tmpArr)
		sarr = sarr[1:]
	}
	t.Log(sarr)
}
