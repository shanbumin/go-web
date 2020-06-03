package main

import (
	"fmt"
	"net/http"
)

//读取请求首部
//一个Header类型的实例就是一个映射，这个映射的键为字符串 ，而键的值则是由任意多个字符串组成的切片。
//即二维数组，键是二维数组的key，值是二维数组的元素。
func headers(w http.ResponseWriter, r *http.Request) {
	//type Header map[string][]string
	//将map[string][]string自定义成了Header类型，本质就是可变元素的关联数组。当一个数组的元素是一维数组，它自然就是二维数组了。
	h := r.Header
    //获取某个具体的首部值
    //h := r.Header["Accept-Encoding"] //返回的是切片
    // h := r.Header.Get("Accept-Encoding")  //返回的是字符串
	fmt.Fprintln(w, h)
}
// http://127.0.0.1:8080/headers
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
