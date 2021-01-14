package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}


//csv文件是一种以逗号分割单元数据的文件，类似表格，但是很轻量。对于存储一些结构化的数据很有用。golang提供了专门处理csv的库。
//和纯文本文件读写类似，csv文件需要通过os创建一个文件句柄，然后调用相关的csv函数读写数据：

func main() {
	// =====================================creating a CSV file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}
    //传入打开的文件句柄，创建csv写
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		//第一行
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	//======================================== reading a CSV file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//FieldsPerRecord是每条记录的预期字段数,-1表示不确定，是可变的
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

	//1
	//Hello World!
	//Sau Sheong
}
