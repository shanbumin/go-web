package main

import (
	"fmt"
	"net/http"
)



//============================/set_cookie   将cookie存储到客户端

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	//===方式一 个人感觉方式一更逼真一些，因为cookie本来就是头信息的一项
	//Cookie结构的String方法可以返回一个经过序列化处理的cookie
	//Set-Cookie响应首部的值就是这些序列化之后的cookie组成的
     w.Header().Set("Set-Cookie",c1.String())
     w.Header().Add("Set-Cookie",c2.String())
    //===方式二
	//http.SetCookie(w, &c1)
	//http.SetCookie(w, &c2)
}

//=================================/get_cookie  从客户端获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
    //====方式一
    //返回了一个切片
	//cookies := r.Header["Cookie"]
	//fmt.Fprintln(w,cookies)//[first_cookie="Go Web Programming"; second_cookie="Manning Publications Co"]
	//fmt.Fprintln(w,cookies[0]) //first_cookie="Go Web Programming"; second_cookie="Manning Publications Co"

	//===方式二(获取的时候推荐采用第二种方式，因为第一种需要我们去解析切片元素，提取我们要的cookie键值对)
    //获取单个cookie
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	c2, err := r.Cookie("second_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}


	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, c2)

	//同时获取多个cookie
	//cs := r.Cookies()
	//fmt.Fprintln(w, cs)
}



//@todo Expires字段和MaxAge字段
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
