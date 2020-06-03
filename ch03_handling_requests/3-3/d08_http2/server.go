package main
import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)
//任意类型只要实现了ServeHTTP,且参数形式如下，则它就是一个处理器，即MyHandler就是一个处理器
type MyHandler struct{}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, //将自定义的处理器绑定到服务器上
	}
	http2.ConfigureServer(&server, &http2.Server{})
	//server.ListenAndServe()
	server.ListenAndServeTLS("cert.pem","key,pem")
}
