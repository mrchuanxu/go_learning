package ch7_test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

// 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口

var EOF = errors.New("EOF")

func TestFprinf(t *testing.T) {
	// var buf bytes.Buffer
	str := "world"
	fmt.Fprintf(os.Stdout, "hello %s", str)
	// t.Log(buf.String(), i, err)

	advane, token, err := bufio.ScanWords([]byte("hel00o"), true)
	t.Log(advane, token, err)
	//io.LimitReader()
}

// func CountingWriter(w io.Writer) (io.Writer, *int64) {

// }
type Reader interface {
	Read(p []byte) (n int, err error)
}

func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}
