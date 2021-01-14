package main

import (
	"gwp/Chapter_2_Go_ChitChat/chitchat/my"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

//main.go文件中的index处理器函数
func index(writer http.ResponseWriter, request *http.Request) {

	//error_message(writer, request, "Cannot get threads")
	//fmt.Fprintln(writer,"欢迎来到首页额")
    //获取所有的帖子
	threads, err := my.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		//检测数据库中是否存在有效的session
		_, err := session(writer, request)
		if err != nil {
			danger(err)
			//如果未登录，则使用public导航条
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			//登录，则使用private导航条
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
