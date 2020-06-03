package main

import (
	"net/http"
	"html/template"
)

//@todo 动态解析模板
func process(w http.ResponseWriter, r *http.Request) {

	 tmpl:=`<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Go Web Programming</title>
  </head>
  <body>
    {{ . }}
  </body>
</html>`

	t:=template.New("tmpl.html")
	t,_=t.Parse(tmpl)
	t.Execute(w, "Hello World004!")

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}


