package main

import (
	"database/sql"
	"fmt"
	"log"
	//隐式调用 _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)



//定义一个指向sql.DB结构的指针
var Db *sql.DB

//初始化
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
		log.Fatal(err)
	}else{
		fmt.Println("db操作对象创建成功")
	}
}

func  main()  {

}




