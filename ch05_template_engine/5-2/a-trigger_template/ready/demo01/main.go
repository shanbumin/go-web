package main

import (
	"html/template"
	"net/http"
)


func process(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../tmpl.html")
	t.Execute(w, "Hello World001!")

}

//@todo 需要在终端调试，因为这里引入的模板是按照相对位置引入的
func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


