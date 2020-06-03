package main
import (
	"fmt"
	"net/http"
)
type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
//---
type WorldHandler struct{}
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
//---
type  SamHandler struct {}
func (this *SamHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Sam is a good man.")
}

//只有借助多路复用器，多个处理器才能生效。因为这几个处理器是绑定到多路复用器上的。
func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	sam:=SamHandler{}
	//使用默认的DefaultServeMux作为多路复用器
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//将自定义的几个处理器绑定到DefaultServeMux
	//http.Handle函数底层调用的就是DefaultServeMux的方法Handle
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	http.Handle("/sam",&sam)
	server.ListenAndServe()
}
