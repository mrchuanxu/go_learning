package go_var_test

import (
	"strings"
	"testing"
)

// 可变参数即其参数数量是可变的 0或者多个 声明可变参数类型前带上...

func toFullname(names ...string) string {
	return strings.Join(names, " ")
}

func TestToFullname(t *testing.T) {
	t.Fatal(toFullname("hello", "jane", "who"))
}
