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
//原生支持
func main() {

	//1.Exec 执行一个SQL语句(添加)
	rs, err :=engine.Exec("insert into user (name,salt,age,passwd) values(?,?,?,?)","sam","salt",18,123456)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(rs)
	//2.Exec 执行一个SQL语句(修改)
	rs, err = engine.Exec("update user set age = ? where name = ?", 30, "sam")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(rs)
	//3.Query 最原始的也支持SQL语句查询，返回的结果类型为 []map[string][]byte。
	qr, err := engine.Query("select * from user")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\r\n",qr)

	//results, err := engine.Where("id = 1").Query()

	//4.QueryString 返回 []map[string]string
	results, err := engine.QueryString("select * from user")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\r\n",results)
	//results, err := engine.Where("a = 1").QueryString()
	//5.QueryInterface 返回 []map[string]interface{}
	res, err := engine.QueryInterface("select * from user")
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v\r\n",res)
	//results, err := engine.Where("a = 1").QueryInterface()


}
