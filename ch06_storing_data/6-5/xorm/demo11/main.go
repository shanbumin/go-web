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

//Sum 求和函数
func main() {
	user:=User{}


	//1.SELECT sum(age) AS total FROM user
	agesFloat64, err := engine.Sum(&user, "age")
	fmt.Printf("agesFloat64=%+v  err=%+v\r\n",agesFloat64,err)
	//2. SELECT sum(age) AS total FROM user
	agesInt64, err := engine.SumInt(&user, "age")
	fmt.Printf("agesInt64=%+v  err=%+v\r\n",agesInt64,err)

	//3.SELECT sum(age), sum(score) FROM user
	//sumFloat64Slice, err := engine.Sums(&user, "age", "score")
	//fmt.Printf("sumFloat64Slice=%+v  err=%+v\r\n",sumFloat64Slice,err)

    //4.SELECT sum(age), sum(score) FROM user
	//sumInt64Slice, err := engine.SumsInt(&user, "age", "score")


}
