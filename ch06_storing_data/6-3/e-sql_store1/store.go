package main

import (
	"database/sql"
	"fmt"
	"log"
	//隐式调用 _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

//定义一个指向sql.DB结构的指针
var Db *sql.DB

//①初始化
//返回的Db对象只是一个数据库操作的对象，它并不是一个连接。
//go封装了连接池，不会暴露给开发者。当Db对象开始数据库操作的时候，go的连接池才会惰性的建立连接，查询完毕之后又会释放连接，连接会返回到连接池之中。
//更多关于数据库的操作，我们将会在后面的mysql专题介绍。
//parseTime=true 解决golang：unsupported Scan, storing driver.Value type []uint8 into type *time.Time

func init() {
	var err error
	//sql.Open方法接收两个参数，第一个参数是指明数据库驱动类型，第二个则是数据库的连接方式字串。返回一个 *sql.DB的指针对象。
	//构建连接, Data Source Name(dsn)格式是："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true")
	//Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
}

//②创建一个帖子 （方法接收者是*Post）
func (post * Post) Create()(err error){
	statment:="insert into posts (content,author) values (?,?)"
	stmt,err:=Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()

	rs,err:=stmt.Exec(post.Content,post.Author)
	if err!=nil{
		return
	}
	id,err:=rs.LastInsertId()
	if err!=nil{
		return
	}
	post.Id=int(id)
	return
}


//③获取一篇帖子
func GetPost(id int) (post Post, err error) {
	post = Post{}
	statment:="select id, content, author from posts where id = ?"
	stmt, err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return
	}
	return
}

//④获取指定个数的帖子
func Posts(limit int) (posts []Post, err error) {
	statement:= "select id, content, author from posts limit ?"
	stmt,err:=Db.Prepare(statement)
	if err!=nil{
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return
	}
	for rows.Next() {//迭代
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}


//⑤更新帖子
//@todo err is shadowed during return  https://blog.csdn.net/wo198711203217/article/details/60574268
func (post *Post) Update() (err error) {
	statement := "update posts set content = ?, author = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec(post.Content, post.Author,post.Id)
	if err != nil {
		return
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Println("Zero affected",post.Id, post.Content, post.Author)
		}
		return err
	}
	return
}


func main() {

	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	//增
	post.Create()//创建完之后,post便有了post.Id值了
	//查
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	//列表
	posts, _ := Posts(10)
	fmt.Println(posts)
	//改
	readPost.Content = "Sam is a good man."
	readPost.Author = "Sam"
	readPost.Update()

	// Delete the post
	//readPost.Delete()
	//
	//// Get all posts
	//posts, _ = Posts(10)
	//fmt.Println(posts) // []

	// Delete all posts
	// DeleteAll()
}




















// =====================================================Delete a post   删除一个帖子
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

// Delete all posts
func DeleteAll() (err error) {
	_, err = Db.Exec("delete from posts")
	return
}


