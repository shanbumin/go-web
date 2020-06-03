package main

import (
	"net/http"
)

func main() {
	//未做任何配置
	http.ListenAndServe("", nil)
}


//multiplexer   ['mʌltipleksə]  简称  mux 多路复用器
//既然ListenAndServe接受的第二个参数是一个处理器，那么为何它的默认值却是多路复用器DefaultServeMux呢？
//答：这是因为DefaultServeMux多路复用器是ServeMux结构的一个实例，
//而后者也拥有ServeHTTP方法，并且这个方法的签名(传参)与成为处理器的签名完全一致。
//换句话说，DefaultServeMux既是ServeMux结构的实例，也是Handler结构的实例。
//ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器。

