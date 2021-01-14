package main

import (
	"html/template"
	"net/http"
)

//虽然Go语言的一般做法是手动地处理错误，但Go也提供了另外一种机制，专门用于处理分析模板时出现的错误：
//todo Must字样的函数一般都会自己判断错误，如果错了直接panic了。
func process(w http.ResponseWriter, r *http.Request) {

	t:= template.Must(template.ParseFiles("../tmpl.html"))
	t.Execute(w, "Hello World005!")

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}