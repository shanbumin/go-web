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
//Exist 检测记录是否存在

func main() {

	//1.SELECT * FROM record_exist LIMIT 1
	has,err:=engine.Exist(new(User))
	fmt.Printf("has=%v err=%v\r\n",has,err)


	//2.SELECT * FROM record_exist WHERE name = ? LIMIT 1
	has, err = engine.Exist(&User{
		Name: "sam",
	})
	fmt.Printf("has=%v err=%v\r\n",has,err)

	//3.SELECT * FROM record_exist WHERE name = ? LIMIT 1
	has, err = engine.Where("name = ?", "sam").Exist(&User{})
	fmt.Printf("has=%v err=%v\r\n",has,err)

    //4.select * from record_exist where name = ?
	has, err = engine.SQL("select * from user where name = ?", "sam").Exist()
	fmt.Printf("has=%v err=%v\r\n",has,err)

	//5. SELECT * FROM record_exist LIMIT 1
	has, err = engine.Table("user").Exist()
	fmt.Printf("has=%v err=%v\r\n",has,err)

    //6. SELECT * FROM record_exist WHERE name = ? LIMIT 1
	has, err = engine.Table("user").Where("name = ?", "sam7").Exist()
	fmt.Printf("has=%v err=%v\r\n",has,err)

}
