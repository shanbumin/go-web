package main
import (
	"fmt"
	"net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func sam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Shanbumin is a good man.")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//绑定自定义的处理器函数到默认的多路复用器上
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	//上述的HandleFunc底层会将处理器函数转成处理器的，本质还是绑定处理器的
	//效果等价如下的操作:
	samHandle:=http.HandlerFunc(sam) //不是调用http.HandlerFunc方法，而是将sam转换成http.HandlerFunc类型
	http.Handle("/sam",&samHandle)

	server.ListenAndServe()
}
