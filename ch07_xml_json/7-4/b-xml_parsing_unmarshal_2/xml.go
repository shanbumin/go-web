package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*

<?xml version="1.0" encoding="utf-8"?>
<post id="1">
  <content>Hello World!</content>
  <author id="2">Sau Sheong</author>
  <comments>
    <comment id="1">
      <content>Have a great day!</content>
      <author>Adam</author>
    </comment>
    <comment id="2">
      <content>How are you today?</content>
      <author>Betty</author>
    </comment>
  </comments>
</post>

 */
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
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


	fmt.Println(post.XMLName.Local)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author)
	fmt.Println(post.Xml)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[1].Id)
	fmt.Println(post.Comments[1].Content)
	fmt.Println(post.Comments[1].Author)
}
