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


//实现一对多以及多对一关系
type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}
type Comment struct {
	Id      int
	Content string
	Author  string
	//Post    *Post
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}


	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments where post_id = ?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		//comment := Comment{Post: &post}
		comment := Comment{}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}




func main() {
	readPost, _ := GetPost(1)
	fmt.Printf("%+v\r\n",readPost)
	fmt.Printf("%+v\r\n",readPost.Comments)

}
