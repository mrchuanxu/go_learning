package implemet_test

import (
	"testing"
	"time"
)

type IReader interface {
	Read() string
}

type IWriter interface {
	Write() string
}

type WR interface {
	IReader
	IWriter
}

type IFile struct {
}

func (ifile *IFile) Read() string {
	return "read"
}

func (IFile *IFile) Write() string {
	return "write"
}

func NewIFile() WR {
	return &IFile{}
}

func TestImp(t *testing.T) {
	// var w io.Writer
	// w = os.Stdout
	// w.Write([]byte("hahah"))
	t1 := time.Unix(1571412817, 0)
	t2 := t1.AddDate(0, 0, 90).AddDate(0, 0, 30).AddDate(0, 0, 62)
	t.Log(t2)
}
