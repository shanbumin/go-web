package main

import (
	"database/sql"
	"fmt"
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



func doSomething(){
	panic("A Panic Running Error")
}

func clearTransaction(tx *sql.Tx){
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil{
		log.Fatalln(err)
	}
}




func main() {

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer clearTransaction(tx)

	rs, err := tx.Exec("UPDATE user SET gold=50 WHERE real_name='vanyarpy'")
	if err != nil {
		log.Fatalln(err)
	}
	rowAffected, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rowAffected)

	rs, err = tx.Exec("UPDATE user SET gold=150 WHERE real_name='noldorpy'")
	if err != nil {
		log.Fatalln(err)
	}
	rowAffected, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rowAffected)

	doSomething()

	if err := tx.Commit(); err != nil {
		// tx.Rollback() 此时处理错误，会忽略doSomthing的异常
		log.Fatalln(err)
	}

}



