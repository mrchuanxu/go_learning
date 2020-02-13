package implemet_test

import (
	"io"
	"os"
	"testing"
)

type IReader interface {
	Read() string
}

type IWriter interface {
	Write() string
	Read() string
}

type IFile struct {
}

func (ifile *IFile) Read() string {
	return "read"
}

func (IFile *IFile) Write() string {
	return "write"
}

func NewIFile() IReader {
	return &IFile{}
}

func TestImp(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hahah"))
}
