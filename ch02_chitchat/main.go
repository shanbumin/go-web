package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets 服务静态文件设置
	mux := http.NewServeMux() //主动创建一个默认的多路复用器，如果不主动，默认底层也会自动帮忙创建的，没有什么不同的，只是前者可以直接调用mux.***去绑定，后者是http.***绑定，本质都一样的
	files := http.FileServer(http.Dir(config.Static)) //public
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// all route patterns matched here
	// route handler functions defined in other files

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread) //处理对帖子的评论
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,//这里可以是自定义的处理器，也可以是自定义的多路复用器，不填则为默认的多路复用器DefaultServeMux(本质也是特殊的处理器)
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
