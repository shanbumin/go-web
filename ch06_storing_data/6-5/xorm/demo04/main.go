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
//Insert 插入一条或者多条记录
func main() {
	user:=User{Name:"sam01",Salt:"salt01",Age:18,Passwd:"p001",Created:time.Now(),Updated:time.Now()}
	user1:=User{Name:"sam02",Salt:"salt01",Age:18,Passwd:"p001",Created:time.Now(),Updated:time.Now()}
	user2:=User{Name:"sam03",Salt:"salt01",Age:18,Passwd:"p001",Created:time.Now(),Updated:time.Now()}
	user3:=User{Name:"sam04",Salt:"salt01",Age:18,Passwd:"p001",Created:time.Now(),Updated:time.Now()}
	user4:=User{Name:"sam05",Salt:"salt01",Age:18,Passwd:"p001",Created:time.Now(),Updated:time.Now()}

	users:=make([]User,0)
	users =append(users,user3)
	users =append(users,user4)


	// INSERT INTO struct () values ()
	engine.Insert(&user)

	// INSERT INTO struct1 () values ()
	// INSERT INTO struct2 () values ()
	engine.Insert(&user1, &user2)
	// INSERT INTO struct () values (),(),()
	engine.Insert(&users)



	// INSERT INTO struct1 () values ()
	// INSERT INTO struct2 () values (),(),()
	//engine.Insert(&user1, &users)


}
