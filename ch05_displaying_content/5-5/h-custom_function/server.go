package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

//@todo 如何传递一个自定义的函数在模板中使用
func process(w http.ResponseWriter, r *http.Request) {

	//创建一个名为FuncMap的映射，并将映射的键设置为函数的名字，而映射的值则设置为实际定义的函数：
	funcMap := template.FuncMap{"fdate": formatDate}
	//将FuncMap与模板进行绑定
	t := template.New("tmpl.html").Funcs(funcMap) //连贯操作
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}



func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
