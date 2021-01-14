package main

import "fmt"

//无论存在什么地方，处理数据的时候都需要把数据读入内存。如果直接存在内存中，不就可以可以直接读了么？的确，数据可以存在内存中。
//所谓内存存储，即定义一些数据结构，数组切片，图或者其他自定义结构，
//把需要持久化的数据存储在这些数据结构中。使用数据的时候可以直接操作这些结构。
//使用映射map作为结构容器的例子
type Post struct {
	Id      int
	Content string
	Author  string
}
//我们定义了两个map的结构PostById，PostByAuthor
//store方法会把post数据存入这两个map结构中。当需要数据的时候，再从这两个内存结构读取即可。
//内存持久化比较简单，严格来说这也不算是持久化，程序退出会清空内存，所保存的数据也会消失。这种持久化只是相对程序运行时而言。
//想要程序退出重启还能读取所存储的数据，这时就得依赖文件或者数据库（非内存数据库）。

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post //这个value是个切片额,因为每个作者可能有多个帖子
func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)  //PostsByAuthor[post.Author]这个值是切片，所以需要追加
}
func main() {
	//初始化上述的两个map，否则无法真正使用
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)
    //初始化几个帖子实例
	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}
	//存储
	store(post1)
	store(post2)
	store(post3)
	store(post4)
	//打印
	fmt.Printf("%+v\r\n",PostById)
	fmt.Printf("%+v\r\n",PostById[1])
	fmt.Printf("%+v\r\n",PostsByAuthor)
	fmt.Printf("%+v\r\n",PostsByAuthor["Pierre"])
}
