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

//Update 更新数据，除非使用Cols,AllCols函数指明，默认只更新非空和非0的字段
//@todo 如果user未指明任何字段，则不作任何更新，默认是会更新updated字段的额
func main() {
	//1.UPDATE user SET ... Where id = ?
	user:=User{Age:21}
	affected, err :=engine.ID(1).Update(&user)
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

	//2.UPDATE user SET ... Where name = ?
	affected, err = engine.Update(&user, &User{Name:"sam"})
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

	//3.UPDATE user SET ... Where id IN (?, ?, ?)
	var ids = []int64{1, 2, 3,7}
	affected, err = engine.In("id",ids).Update(&user)
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

	//4.通过Cols强制更新指示的列 UPDATE user SET age = ?, updated=? Where id = ?
	affected, err = engine.ID(1).Cols("age").Update(&User{Name:"sam", Age: 12})
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

	//5.f通过省略不强制更新指示的列   UPDATE user SET age = ?, updated=? Where id = ?
	affected, err = engine.ID(1).Omit("name").Update(&User{Name:"rick", Age: 13})
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

	//6.UPDATE user SET name=?,age=?,salt=?,passwd=?,updated=? Where id = ?
	//@todo 强制更改所有的列，结构体实例未设置，则将修改为空了
	affected, err = engine.ID(1).AllCols().Update(&user)
	fmt.Printf("affected=%+v  err=%+v\r\n",affected,err)

}
