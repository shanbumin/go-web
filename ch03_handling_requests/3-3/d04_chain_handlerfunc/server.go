package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)
//串联两个处理器函数
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
//传递一个处理器函数返回一个处理器函数
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//反射获取h的名字，即hello
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(w)
		fmt.Println(r)
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
//当f1作为参数传递给f2的时候，把这种情形称为f1与f2串联。
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello)) //串联处理器函数
	server.ListenAndServe()
}