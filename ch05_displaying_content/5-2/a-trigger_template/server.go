package main

import (
	"html/template"
	"net/http"
)


//==================== /process



func process(w http.ResponseWriter, r *http.Request) {
	//@todo 下面两句话发生的事情还是蛮惊人的，看书中所说即可
	//多模板文件进行语法分析
	//返回一个Template类型的已分析模板和一个错误作为结果
	t, _ := template.ParseFiles("tmpl.html")
	//将数据应用到模板里面
	t.Execute(w, "Hello World!")
}
//http://127.0.0.1:8080/process
func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


