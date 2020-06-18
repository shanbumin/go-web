package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//定义一些结构体用于表示数据
//xml的命名空间和元素名 https://www.runoob.com/xml/xml-namespaces.html

/*
<post id="1">
  <content>Hello World!</content>
  <author id="2">shanbumin</author>
</post>
 */

//go语言使用标签来决定如何对结构以及XML元素进行映射
type Post struct {
	XMLName xml.Name `xml:"post"`   //特殊①

	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`

	Xml     string   `xml:",innerxml"` //特殊②
}

type Author struct {
	Id   string `xml:"id,attr"`  //这样的结构化标签会将 <author id="2">shanbumin</author>中的属性id的值提取出来的
	Name string `xml:",chardata"`// 这样的结构化标签会将<author id="2">shanbumin</author>中的shanbumin字符提取出来的
}


//todo xml转结构体
func main() {

    //第一步:打开xml文件
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

    //第二步:读取xml内容
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	//第三步:将xml数据解封到结构里面
	var post Post
	xml.Unmarshal(xmlData, &post)
	//fmt.Println(post)
	fmt.Println("XMLName.Space",post.XMLName.Space)
	fmt.Println("XMLName.Local",post.XMLName.Local)
	fmt.Println("Id",post.Id)
	fmt.Println("Content",post.Content)
	fmt.Println("Author",post.Author)
	fmt.Println("Author.Id",post.Author.Id)
	fmt.Println("Author.Name",post.Author.Name)
	fmt.Println("Xml",post.Xml)
}
