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

//上下文缓存，如果启用，那么针对单个对象的查询将会被缓存到系统中，可以被下一个查询使用。
func main() {

	//sess := engine.NewSession()
	//defer sess.Close()
	//
	//var context = xorm.NewMemoryContextCache()
	//
	//var c2 ContextGetStruct
	//has, err := sess.ID(1).ContextCache(context).Get(&c2)
	//assert.NoError(t, err)
	//assert.True(t, has)
	//assert.EqualValues(t, 1, c2.Id)
	//assert.EqualValues(t, "1", c2.Name)
	//sql, args := sess.LastSQL()
	//assert.True(t, len(sql) > 0)
	//assert.True(t, len(args) > 0)
	//
	//var c3 ContextGetStruct
	//has, err = sess.ID(1).ContextCache(context).Get(&c3)
	//assert.NoError(t, err)
	//assert.True(t, has)
	//assert.EqualValues(t, 1, c3.Id)
	//assert.EqualValues(t, "1", c3.Name)
	//sql, args = sess.LastSQL()
	//assert.True(t, len(sql) == 0)
	//assert.True(t, len(args) == 0)




}
