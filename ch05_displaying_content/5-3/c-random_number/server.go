package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err:= template.ParseFiles("tmpl.html")
	if err != nil {
		//找不到的原因是你的IDE默认配置的Working Directory是到go-web下的，
		//所以建议你在控制台执行go run 或者配置一下Working Directory
		fmt.Println(err)
		return
	}
	//随机的生成介于0至10之间的随机整数，然后通过判断这个随机整数是否大于5来创建出一个布尔值，并在最后将这个布尔值传递给模板
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
