package web

import (
	"testing"
	"net/http"
	"time"
)

func TestWeb(t *testing.T) {
	test()
}

func hello(w http.ResponseWriter, r *http.Request) {
	//获取body
	//bytes:=make([]byte,r.ContentLength)
	//r.Body.Read(bytes)
	//fmt.Println(string(bytes))

	/*a:=r.Form["id"]
	for v :=range a{
		fmt.Println(v)
	}*/

}

func test() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func web1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	//http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func web2() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":4000", mux)
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func web3() {
	server := http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/", hello)
	server.ListenAndServe()
}
