package main

import (
	"fmt"
	"net/http"
)
//首先使用了ParseForm方法对请求进行语法分析，然后再访问Form字段，获取具体的表单
//注意r.Form   r.PostForm   r.MultipartForm的值到底是多少还取决于客户端表单的编码
//fmt.Fprintln(w, r.PostForm)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w,"Form字段",r.Form,"PostForm字段",r.PostForm,"MultipartForm字段",r.MultipartForm)
	//fmt.Fprintln(w,r.Form,r.Form["hello"],r.Form["hello"][0],r.Form["hello"][1])
}
//Form的感觉就是不管是get还是post请求，都会获取，类似php的$_REQUEST
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
