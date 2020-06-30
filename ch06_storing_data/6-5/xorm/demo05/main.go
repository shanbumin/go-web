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
//Get 查询单条记录

func main() {

	//1.SELECT * FROM user LIMIT 1
	user:=User{}
	engine.Get(&user)
	fmt.Printf("%+v\r\n",user)

	//2.SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
	user=User{}
	engine.Where("name = ?","sam").Desc("id").Get(&user)
	fmt.Printf("%+v\r\n",user)

	//3.SELECT name FROM user WHERE id = ?
	var name string
	engine.Table(&user).Where("id = ?",1).Cols("name").Get(&name)
	fmt.Printf("%+v\r\n",name)

	//4.SELECT id FROM user WHERE name = ?
	var id int64
	engine.Table(&user).Where("name = ?", name).Cols("id").Get(&id)
	fmt.Printf("%+v\r\n",id)
	engine.SQL("select id from user order by id desc").Get(&id)
	fmt.Printf("%+v\r\n",id)
	//5.SELECT * FROM user WHERE id = ?
	var valuesMap = make(map[string]string)
	engine.Table(&user).Where("id = ?", id).Get(&valuesMap)
	fmt.Printf("%+v\r\n",valuesMap)

	//6.SELECT col1, col2, col3 FROM user WHERE id = ?
	cols:=make([]string,3)
	cols[0]="id"
	cols[1]="name"
	cols[2]="age"
	var valuesSlice = make([]interface{}, len(cols))
	engine.Table(&user).Where("id = ?", id).Cols(cols...).Get(&valuesSlice)
	fmt.Printf("%+v\r\n",valuesSlice)



}
