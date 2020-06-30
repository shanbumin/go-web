package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)


type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}





var engine *xorm.Engine
var err error

func init(){
	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	//第一步创建引擎
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
		log.Fatal(err)
	}
}

//在一个Go程中多次操作数据库，但没有事务
func main() {
	session := engine.NewSession()
	defer session.Close()

	user1 := User{Name: "xiaoxiao",Age:10,Created: time.Now()}
	if _, err := session.Insert(&user1); err != nil {
		log.Fatal(err)
	}

	user2 := User{Name: "yy",Age:100,Created: time.Now()}
	if _, err := session.Where("id = ?", 2).Update(&user2); err != nil {
		log.Fatal(err)
	}
	if _, err := session.Exec("delete from user where name = ?", user2.Name); err != nil {
		log.Fatal(err)
	}



}
