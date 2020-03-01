package ch8_test

import (
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var testSuku = [][]byte{
	{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
	{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
	{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
	{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
	{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
	{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
	{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}

// 和map类似，channel也对应一个make创建的底层数据结构的引用。

func TestRoutine(t *testing.T) {
	ch := make(chan int)
	// if ch == nil {
	// 	err := errors.New("ch has nothing")
	// 	panic(err)
	// }
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
	go func() {
		for ival := range ch {
			t.Log(ival)
		}
	}()
	t.Log("length", len(ch))
	time.Sleep(time.Second * 2)
	close(ch)
}

func TestDial(t *testing.T) {
	conn, err := net.DialTCP("tcp", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	// mustCopy(conn, os.Stdin)
	conn.CloseRead()
	<-done // wait for background goroutine to finish
}

func TestSuku(t *testing.T) {
	ret := isValidSudoku(testSuku)
	t.Log(ret)
}

func isValidSudoku(board [][]byte) bool {
	// 先判断行列
	// 再判断 内部9个
	if len(board) == 0 {
		return false
	}

	for i := 0; i < 9; i++ {
		colMap := make(map[byte]bool)
		lineMap := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if _, ok := colMap[board[i][j]]; ok && board[i][j] != '.' {
				return false
			}
			colMap[board[i][j]] = true
			if _, ok := lineMap[board[j][i]]; ok && board[j][i] != '.' {
				return false
			}
			lineMap[board[j][i]] = true
		}
	}
	// 判断内旋
	rets := true
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !(rets && isValidSubSudoku(board, i, i+3, j, j+3)) {
				return false
			}
		}
	}
	return rets
}

func isValidSubSudoku(board [][]byte, istart, iend, jstart, jend int) bool {
	validMap := make(map[byte]bool)
	for istart < iend {
		for j := jstart; j < jend; j++ {
			if _, ok := validMap[board[istart][j]]; ok && board[istart][j] != '.' {
				return false
			}
			validMap[board[istart][j]] = true
		}
		istart++
	}
	return true
}
