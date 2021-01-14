package main

import (
	"html/template"
	"net/http"
)

//为了展示模板中的上下文感知特性而设置的处理器
//@todo 本质就是自动转义或者说过滤吧?
//@todo curl  -i   127.0.0.1:8080/process

/*

<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Go Web Programming</title>
  </head>
  <body>
    <div>I asked: &lt;i&gt;&#34;What&#39;s up?&#34;&lt;/i&gt;</div>
    <div><a href="/I%20asked:%20%3ci%3e%22What%27s%20up?%22%3c/i%3e">Path</a></div>
    <div><a href="/?q=I%20asked%3a%20%3ci%3e%22What%27s%20up%3f%22%3c%2fi%3e">Query</a></div>
    <div><a onclick="f('I asked: \x3ci\x3e\x22What\x27s up?\x22\x3c\/i\x3e')">Onclick</a></div>
  </body>
</html>

 */

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
