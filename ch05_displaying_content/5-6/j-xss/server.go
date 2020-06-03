package main

import (
  "net/http"
  "html/template"
)


//访问并输入
//http://127.0.0.1:8080/xss/form
//<script>alert("sam is a good man");</script>

func process(w http.ResponseWriter, r *http.Request) {
  //为了演示，可以强制暂时要浏览器关闭内置的XSS防御功能
  w.Header().Set("X-XSS-Protection", "0")
  t, _ := template.ParseFiles("tmpl.html")  
  //t.Execute(w, r.FormValue("comment"))

  //@todo 强制不转义,默认模板引擎是上下文感知的，会自动转义的
  t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {  
  t, _ := template.ParseFiles("form.html")  
  t.Execute(w, nil)  
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/", form)
  server.ListenAndServe()
}
