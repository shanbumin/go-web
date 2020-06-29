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

//更新帖子
//@todo err is shadowed during return  https://blog.csdn.net/wo198711203217/article/details/60574268
func (post *Post) Update() (err error) {
	//1.Db.Prepare
	stmt, err := Db.Prepare("update posts set content = ?, author = ? where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	//2.stmt.Exec
	ret, err := stmt.Exec(post.Content, post.Author,post.Id)
	if err != nil {
		return
	}
	//3.影响行数
	_, err = ret.RowsAffected()
	if err !=nil{
		return
	}
	//fmt.Println(rf,err)
	//1 <nil>
	//0 <nil>
	return
}
func main() {
	readPost:=&Post{}
	//改
	readPost.Id=1
	readPost.Content = "content"
	readPost.Author = "author"
	err:=readPost.Update()
	if err !=nil{
		fmt.Println(err)
	}
}


