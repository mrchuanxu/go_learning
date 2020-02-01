package ch5_test

import (
	"os"
	"testing"
)

// 首先创建一些目录，然后将其删除
func TestDirCreate(t *testing.T) {
	var rmDirs []func()
	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmDirs = append(rmDirs, func() {
			os.RemoveAll(dir) // 这里是同一个地址
		})
	}
	for _, rmdir := range rmDirs {
		rmdir()
	}
}

func tempDirs() []string {
	dirs := []string{"a.txt", "b.txt", "c.txt", "d.txt", "e.txt"}
	return dirs
}
