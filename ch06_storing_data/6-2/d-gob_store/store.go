package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

//===========================store data
func store(data interface{}, filename string) {
	//申请一块缓冲(可以看成是申请一块内存啦，写过c的人都知道)
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	//将缓冲中的内容写到文件中
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

//===========================================load the data
func load(data interface{}, filename string) {
	//读取出内容，返回的是字节切片
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//将读取的内容放入缓冲中
	buffer := bytes.NewBuffer(raw)
	//将缓冲中的内容使用gob解码即可
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}

	//存储二进制流
	store(post, "post1")
	//加载
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
