package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//io/ioutil包  =====>ioutil包提供给外部使用的一共有1个变量，7个方法。
//https://www.jianshu.com/p/6b7ababfcced
//http://www.cnblogs.com/golove/p/3278444.html

func process(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1024)

	//从MultipartForm字段的File字段里面取出文件头FileHeader
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		// ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
		// 如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
		// 所有数据，所以不会把 EOF 当做错误处理。
		//func ReadAll(r io.Reader) ([]byte, error)
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
