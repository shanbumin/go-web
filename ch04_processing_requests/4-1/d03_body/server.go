package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	//先获取主体数据的字节长度
	len := r.ContentLength
	//根据这个长度创建一个字节切片
	//注意这个长度非常之重要，因为Read方法会自动读取这个长度的字节并放到body变量中的
	body := make([]byte, len)
	//调用Read方法将主体数据读取到字节数组中
	r.Body.Read(body)//
	fmt.Fprintln(w, string(body))
}
//curl  -id   "first_name=sausheong&last_name=chang"    127.0.0.1:8080/body
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
