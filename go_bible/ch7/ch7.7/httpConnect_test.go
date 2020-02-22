package transconnect

import (
	"fmt"
	"net/http"
	"testing"
)

// Handler 操作
type Handler interface {
	TransServeHTTP(w ReponseWriter, r *Request)
}

// ReponseWriter 返回
type ReponseWriter struct {
}

// Request 请求
type Request struct {
}

// ListenAndServe 监听
func ListenAndServe(address string, h Handler) error {
	return nil
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// TestTransHttp 测试trans的http服务
func TestTransHttp(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	t.Fatal(http.ListenAndServe("localhost:8000", mux))
	// nil 使用DefaultServeMux作为服务器
	// db := database{"shoes": 50, "socks": 5}
	// http.HandleFunc("/list", db.list)
	// http.HandleFunc("/price", db.price)
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
