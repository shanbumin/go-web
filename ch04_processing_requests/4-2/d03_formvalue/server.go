package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "FormValue",r.FormValue("hello"))
	fmt.Fprintln(w, "PostFormValue",r.PostFormValue("hello"))
	fmt.Fprintln(w, "Form",r.Form)
	fmt.Fprintln(w, "PostForm",r.PostForm)
	fmt.Fprintln(w, "MultipartForm",r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
