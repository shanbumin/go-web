package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var db *sql.DB

func init(){
	var err error
	//创建一个数据库抽象对象
	db, err= sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil{
		log.Fatal(err)
	}

}


func  query001(){

	//下面两种方式的底层通信不完全一样。一种你是plaintext方式，另外一种是prepared方式。
	rows, err := db.Query("SELECT real_name FROM user WHERE gid = 1000")
	//rows, err := db.Query("SELECT real_name FROM user WHERE gid = ?", 1000)
	if err != nil{
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next(){
		var s string
		err = rows.Scan(&s)
		if err !=nil{
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}


}

func  query002(){

}






func main() {

	defer db.Close()

	query001()








}



