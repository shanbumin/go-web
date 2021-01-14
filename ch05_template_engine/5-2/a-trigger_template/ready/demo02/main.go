package main

import (
	"html/template"
	"net/http"
)


func process(w http.ResponseWriter, r *http.Request) {

	t:=template.New("tmpl.html")
	t,_=t.ParseFiles("../tmpl.html")
	t.Execute(w, "Hello World002!")


}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


