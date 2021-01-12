package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("./cert.pem", "./key.pem")
}


//1.搞清楚SSL证书         cert.pem    X.509证书
//2.搞清楚服务器的私钥     key.pem
//3.搞清楚SSL证书与服务器私钥之间的关系
//4.自制的SSL证书  <=====>服务器签名
//  CA机构颁发的SSL证书<=====>服务器签名+CA签名


