package main

import (
	"fmt"
	"net/url"
)

func main() {

	u, err := url.Parse("http://root:root@bing.com:8080/dc/search?q=dotnet&name=shanbumin#first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n",u) //*url.URL

	fmt.Println("Scheme:",u.Scheme) //http
	fmt.Println("Opaque:",u.Opaque)// 编码后的不透明数据
	fmt.Println("User:",u.User)// 用户名和密码信息  root:root
	fmt.Println("Host:",u.Host) // host或host:port    bing.com:8080
	fmt.Println("Path:",u.Path) //  /dc/search
	fmt.Println("RawQuery:",u.RawQuery)// 编码后的查询字符串，没有'?'   q=dotnet&name=shanbumin
	fmt.Println("Fragment:",u.Fragment)// 引用的片段（文档位置），没有'#'   first
	//u.Scheme = "https"
	//u.Host = "google.com"
	//q := u.Query()
	//q.Set("q", "golang")
	//u.RawQuery = q.Encode()
	//fmt.Println(u)
}
