package main

import (
	"net/http"
	"html/template"
)

//@todo 解析多个模板
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("../*.html")
	t.Execute(w, "Hello World003!")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


