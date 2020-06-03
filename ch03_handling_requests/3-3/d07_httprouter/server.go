package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

//@todo  go get github.com/julienschmidt/httprouter
//http://127.0.0.1:8080/hello/sam
func main() {
	//创建一个自定义的多路复用器
	mux := httprouter.New()
	//这个程序不再使用HandleFunc绑定处理器函数，而是直接把处理器函数与给定的HTTP方法进行绑定
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
