package main

import (
	"fmt"
)
//所谓内存存储，即定义一些数据结构，数组切片，图或者其他自定义结构，
//把需要持久化的数据存储在这些数据结构中。使用数据的时候可以直接操作这些结构。
//使用映射map作为结构容器的例子

type Post struct {
	Id      int
	Content string
	Author  string
}

//我们定义了两个map的结构PostById，PostByAuthor，store方法会把post数据存入这两个map结构中。当需要数据的时候，再从这两个内存结构读取即可。
//内存持久化比较简单，严格来说这也不算是持久化，程序退出会清空内存，所保存的数据也会消失。这种持久化只是相对程序运行时而言。
//想要程序退出重启还能读取所存储的数据，这时就得依赖文件或者数据库（非内存数据库）。

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post //这个value是个切片额


func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	//map[1:0xc0000641e0 2:0xc000064210 3:0xc000064240 4:0xc000064270]
	//map[Pedro:[0xc000064240] Pierre:[0xc000064210] Sau Sheong:[0xc0000641e0 0xc000064270]]

	fmt.Println(PostById)
	fmt.Println(PostsByAuthor)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	//PostsByAuthor["Sau Sheong"]是个切片
	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	//PostsByAuthor["Pedro"]是个切片
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
	//&{1 Hello World! Sau Sheong}
	//&{2 Bonjour Monde! Pierre}
	//&{1 Hello World! Sau Sheong}
	//&{4 Greetings Earthlings! Sau Sheong}
	//&{3 Hola Mundo! Pedro}
}
