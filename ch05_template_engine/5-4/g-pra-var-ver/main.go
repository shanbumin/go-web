package main

import (
	"html/template"
	"net/http"
)


//传递一个map
func process(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("tmpl.html")
	mp:=map[string]interface{}{"name":"sam","age":18,"sex":"male"}
	t.Execute(w, mp)

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


