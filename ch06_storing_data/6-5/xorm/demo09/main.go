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

//Delete 删除记录，需要注意，删除必须至少有一个条件，否则会报错。要清空数据库可以用EmptyTable


func main() {

	user:=User{}
	//1.DELETE FROM user Where ...
	affected, err := engine.Where("name=?","sam").Delete(&user)
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)


	//2. DELETE FROM user Where id = ?
	affected, err = engine.ID(1).Delete(&user)
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)


}
