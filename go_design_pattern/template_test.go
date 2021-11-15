package go_design_pattern_test

import "testing"

// 模版方法 将实现推迟
type ReaderIFace interface {
	Read(file string) string
	ReadAll(file []string) string
}

// 模版方法
func IOTemplate(file string, r ReaderIFace) string {
	// ... 实现逻辑
	r.Read(file)
	r.ReadAll([]string{file})
	return ""
}

type reader struct{}

func (r *reader) Read(file string) string {
	return file
}
func (r *reader) ReadAll(file []string) string {
	return "all"
}

func Test_template(t *testing.T) {
	IOTemplate("txt", &reader{})
}

// callback 回调 A调用B B反过来调用A
