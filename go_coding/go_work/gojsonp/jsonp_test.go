package gojsonp

import (
	"testing"
	"github.com/thedevsaddam/gojsonq"
	"time"
)

func TestJsonp(t *testing.T){
	const jsonStr = `{"movies":[{"name":"Pirates of the Caribbean","sequels":[{"name":"The Curse of the Black Pearl","released":2003},{"name":"Dead Men Tell No Tales","released":2017}]}]}`
	jq := gojsonq.New().JSONString(jsonStr)

	name := jq.Find("movies.[0].name")
	t.Log("name:",name)
	fsequel := jq.Reset().Find("movies.[0].sequels.[0].name")
	t.Log("first sequel: ",fsequel)

	t1 := time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Millisecond*2)
	t2 := time.Now().Format("2006-01-02 15:04:05")
	t.Log(t2-t1)
	
}