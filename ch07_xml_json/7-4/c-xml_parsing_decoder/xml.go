package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

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
	//第一步：打开文件
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	//第二步：根据给定的XML数据生成相应的解码器
	decoder := xml.NewDecoder(xmlFile)
	//第三步：每迭代一次解码器中的所有XML数据
	for {
		//每进行一次迭代，就从解码器里面获取一个token
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
        //检查token的类型
		switch se := t.(type) {
		case xml.StartElement:
			fmt.Println(se.Name.Local)
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)//将xml数据解码至结构
				fmt.Println(comment)
			}
		}
	}
}
