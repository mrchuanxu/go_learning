package ch1

import (
	"io"
	"os"
	"testing"
)

func ReadFrom(reader io.Reader,num int)([]byte,error){
	p := make([]byte,num)
	n,err := reader.Read(p)
	if n>0{
		return p[:n],nil
	}
	return p,err
}

func TestIO(t *testing.T){
	data,err  := ReadFrom(os.Stdin,11)
	t.Log(data,err)
}