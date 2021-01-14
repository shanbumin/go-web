package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//测试handleGet处理器函数
func TestHandleGet(t *testing.T) {
	//创建一个用于运行测试的多路复用器
	mux := http.NewServeMux()
	//绑定想要测试的处理器
	mux.HandleFunc("/post/", handleRequest)
	//************************************
	writer := httptest.NewRecorder() //创建记录器，用于获取服务器返回的HTTP响应
	request, _ := http.NewRequest("GET", "/post/1", nil)//为被测试的处理器创建相应的请求
	mux.ServeHTTP(writer, request)//把创建出的记录器以及HTTP请求传递给多路复用器
	//*****************************************************

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

//测试handlePut处理器函数
func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"putsir"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
