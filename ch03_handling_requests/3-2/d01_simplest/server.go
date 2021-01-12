package main

import (
	"log"
	"net/http"
)


//既然ListenAndServe接受的第二个参数是一个处理器，那么为何它的默认值却是多路复用器DefaultServeMux呢？
func main() {
	//未做任何配置
	err:=http.ListenAndServe("", nil)
	if err!=nil{
		log.Fatal(err)
	}

}

