package gojsonp

import (
	"testing"
	// "github.com/thedevsaddam/gojsonq"

	"time"
)

func squares(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func TestJsonp(t *testing.T) {
	sum := 0
	for term := range squares(4) {
		sum += term
	}
	t.Log(sum)
	// const jsonStr = `{"movies":[{"name":"Pirates of the Caribbean","sequels":[{"name":"The Curse of the Black Pearl","released":2003},{"name":"Dead Men Tell No Tales","released":2017}]}]}`
	// jq := gojsonq.New().JSONString(jsonStr)

	// name := jq.Find("movies.[0].name")
	// t.Log("name:",name)
	// fsequel := jq.Reset().Find("movies.[0].sequels.[0].name")
	// t.Log("first sequel: ",fsequel)

	// t1 := time.Now().Format("2006-01-02 15:04:05")
	// time.Sleep(time.Millisecond*2)
	// t2 := time.Now().Format("2006-01-02 15:04:05")
	tim := time.Now().AddDate(0, 0, -5)
	// zeroTim := time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location()).Unix()
	// t.Log(zeroTim)
	// t.Log(reflect.TypeOf(zeroTim))
	t.Log(tim)
}
