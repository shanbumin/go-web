package main

import (
	"fmt"
	"net/http"
)

//======================自定义一个处理器  MyHandler=================================
//任意类型只要实现了Handler接口，即只需要实现ServeHTTP方法,且参数形式如下，则它就是一个处理器，即MyHandler就是一个处理器
type MyHandler struct{}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
//================================================================
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, //将自定义的处理器绑定到服务器上
	}
	server.ListenAndServe()
}
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务。
//ServeHTTP应该将回复的头域和数据写入ResponseWriter接口然后返回。
//返回标志着该请求已经结束，HTTP服务端可以转移向该连接上的下一个请求。
