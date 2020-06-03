package main

import (
	"fmt"
	"net/http"
)

//ParseMultipartForm将请求的主体作为multipart/form-data解析。
//请求的整个主体都会被解析，得到的文件记录最多maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。
//如果必要，ParseMultipartForm会自行调用ParseForm。重复调用本方法是无意义的。

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w,"Form字段",r.Form,"PostForm字段",r.PostForm,"MultipartForm字段",r.MultipartForm)
	//fmt.Fprintln(w, r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
