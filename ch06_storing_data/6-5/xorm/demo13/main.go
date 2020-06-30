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

//迭代
//Iterate 和 Rows 根据条件遍历数据库，可以有两种方式: Iterate and Rows
func main() {

	//1.SELECT * FROM user
	 engine.Iterate(&User{Name:"sam01"}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Printf("%+v\r\n",user)
		return nil
	})
	 fmt.Println("-------------------------------------------------------------------")
	 //2.SELECT * FROM user Limit 0,2
	    //SELECT * FROM user Limit 2, 4
	    //...todo 不断的这样迭代2个下去，直到取完
	 engine.BufferSize(2).Iterate(&User{Name:"sam01"}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		 fmt.Printf("%+v\r\n",user)
		return nil
	})


	fmt.Println("-------------------------------------------------------------------")
	//3.SELECT * FROM user
	rows,_:= engine.Rows(&User{Name:"sam01"})
	defer rows.Close()

	for rows.Next() {
		user:=User{}
		err = rows.Scan(&user)
		fmt.Printf("%+v\r\n",user)
	}


}
