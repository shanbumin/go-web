package main

import (
	"net/http"
)

func main() {
	//Server结构体变量
	//通过Server结构体调用可以配置很多东西额
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
