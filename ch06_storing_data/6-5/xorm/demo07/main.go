package main

import (
	"fmt"
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

type Detail struct {
	Id int64
	UserId int64 `xorm:"index"`
}

type UserDetail struct {
	User `xorm:"extends"`
	Detail `xorm:"extends"`
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


//Find 查询多条记录，当然可以使用Join和extends来组合使用
func main() {

	var users []User
	//1.SELECT * FROM user WHERE name = ? AND age > 10 limit 10 offset 0
	engine.Where("name = ?","sam").And("age >= 30").Limit(10, 0).Find(&users)
	fmt.Printf("%+v\r\n",users)

	//2.SELECT user.*, detail.* FROM user INNER JOIN detail WHERE user.name = ? limit 10 offset 0
	var users2 []UserDetail
	engine.Table("user").Select("user.*, detail.*").Join("INNER", "detail", "detail.user_id = user.id").Where("user.name = ?","sam").Limit(10, 0).Find(&users2)
	fmt.Printf("%+v\r\n",users2)

}
