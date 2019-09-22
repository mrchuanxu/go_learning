package main

import (
	"path/filepath"
	"os"
	"os/exec"
	"strings"
	"fmt"
)

func GetAppPath() string{
	file,_ := exec.LookPath(os.Args[0])
	path,_ := filepath.Abs(file)
	index := strings.LastIndex(path,string(os.PathSeparator))

	return path[:index]
}

func main(){
	fmt.Println(GetAppPath())
}