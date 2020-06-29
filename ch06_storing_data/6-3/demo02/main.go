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


//创建一个帖子 （方法接收者是*Post）
func (post * Post) Create()(err error){

	//1.Db.Prepare
	stmt,err:=Db.Prepare("insert into posts (content,author) values (?,?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	//2.stmt.Exec
	rs,err:=stmt.Exec(post.Content,post.Author)
	if err!=nil{
		return
	}
	//3.返回值
	id,err:=rs.LastInsertId()
	if err!=nil{
		return
	}
	post.Id=int(id)
	return
}


func main() {

	post := Post{Content: "Hello World!", Author: "sam"}
	err:=post.Create()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v",post)

}


