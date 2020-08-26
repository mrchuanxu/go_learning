package nethttp_test

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_First(t *testing.T) {
	resp, _ := http.Get("http://cn.bing.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}

func Test_Handle(t *testing.T) {
	s := &http.Server{
		Addr: ":8090",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
		},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
