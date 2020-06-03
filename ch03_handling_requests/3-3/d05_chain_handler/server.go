package main

import (
	"fmt"
	"net/http"
)
//==============================================
type HelloHandler struct{}
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
//===============================================


//串联处理器,在真正执行处理器之前，需要调用几个方法，先调用protect，再调用log
//log方法  传递一个处理器返回处理器接口类型
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

//protect方法
func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// some code to make sure the user is authorized
		fmt.Println("protect")
		h.ServeHTTP(w, r)
	})
}

//串联处理器
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
