package main

import (
	"html/template"
	"net/http"
)


func process(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../t1.html","../t2.html")
	t.Execute(w, "Hello World!") //默认只会展示t1,传递的值也是给t1的

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

