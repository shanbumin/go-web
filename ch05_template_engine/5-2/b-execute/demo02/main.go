package main

import (
	"html/template"
	"net/http"
)


func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../t1.html","../t2.html")
	t.ExecuteTemplate(w,"t2.html","Hello World!") //指明返回t2，传递的值返回给t2
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}



