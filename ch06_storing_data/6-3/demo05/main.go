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

//删除一个帖子
func (post *Post) Delete() (err error) {
	//1.Db.Prepare
	stmt, err := Db.Prepare("delete from posts where id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()
	//2.stmt.Exec
	res, err := stmt.Exec(post.Id)
	num, err := res.RowsAffected()
	fmt.Println(num,err)
	return
}
//删除所有的帖子
func (post *Post) DeleteAll() (err error) {
	//1.Db.Prepare
	stmt, err := Db.Prepare("delete from posts")
	if err != nil {
		return
	}
	defer stmt.Close()
	//2.stmt.Exec
	res, err := stmt.Exec()
	num, err := res.RowsAffected()
	fmt.Println(num,err)
	return

}

func main() {
	readPost:=&Post{}
	//删一条
	readPost.Id=1
	err:=readPost.Delete()
	if err !=nil{
		fmt.Println(err)
	}
	//删除所有
	err=readPost.DeleteAll()
	if err !=nil{
		fmt.Println(err)
	}
}


