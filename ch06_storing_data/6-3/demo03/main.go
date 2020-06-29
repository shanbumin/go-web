package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



var Db *sql.DB
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

//---------------------  Post --------------------------------------------------
type Post struct {
	Id      int
	Content string
	Author  string
}


//根据主键获取一篇帖子
func GetPostById(id int) (post Post, err error) {
	post = Post{}
	//1.Db.Prepare
	stmt, err := Db.Prepare("select id, content, author from posts where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	//2.stmt.QueryRow
	err = stmt.QueryRow(id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return
	}
	return
}


//获取指定个数的帖子
func Posts(limit int) (posts []Post, err error) {
	//1.Db.Prepare
	stmt,err:=Db.Prepare("select id, content, author from posts limit ?")
	if err!=nil{
		return
	}
	defer stmt.Close()
	//2.stmt.Query(
	rows, err := stmt.Query(limit)
	if err != nil {
		return
	}
	//3.迭代
	for rows.Next() {
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



func main() {
	//查
	readPost, _ := GetPostById(1)
	fmt.Printf("%+v\r\n",readPost)
	//查询多个
	posts,_:=Posts(3)
	fmt.Printf("%+v",posts)
}


