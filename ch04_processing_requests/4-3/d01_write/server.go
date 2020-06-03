package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

//===============================/write
//http://127.0.0.1:8080/write
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

//===============================/writeheader
//http://127.0.0.1:8080/writeheader
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}
//==================================/redirect
//http://127.0.0.1:8080/redirect
func headerExample(w http.ResponseWriter, r *http.Request) {
	//通过编写首部实现客户端重定向
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(302)//调用这个方法之后不能对响应首部做任何写入操作,所以对首部的设置一定要在设置状态码之前进行
}

//=================================== /json
//http://127.0.0.1:8080/json
func jsonExample(w http.ResponseWriter, r *http.Request) {
	//直接向客户端返回json数据
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

//curl  -i   127.0.0.1:8080/write
//curl  -i   127.0.0.1:8080/writeheader
//curl  -i   127.0.0.1:8080/redirect
//curl  -i   127.0.0.1:8080/json

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
