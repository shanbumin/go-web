package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//分析模板的方法不变
	t, _ := template.ParseFiles("layout.html")
	//但是这次在执行模板的时候，程序需要显式地使用ExecuteTemplate方法，并把待执行的layout模板的名字用作方法的第二个参数。
	//因为layout模板嵌套了content模板，所以程序只需要执行layout模板就可以在浏览器中得到content模板产生的输出了
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
