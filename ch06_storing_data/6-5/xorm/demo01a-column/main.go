package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)



type Student struct {
	Id        int64 `xorm:"pk autoincr"` // 注：使用getOne 或者ID() 需要设置主键
	Age      int
	Name     string
	ImageUrl  string `xorm:"'img_url'"`  //使用 xorm:"'column_name'"可以使该 field 对应的 Column 名称为指定名称,注意里面的单引号额
	Url       string
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `xorm:"'deleted'"` // 此特性会激发软删除
}

func (m *Student) TableName() string {
	return "student1"
}

func main() {

	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	//第一步创建引擎
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
      log.Fatal(err)
	}
	//第二步 同步创建表
	err=engine.Sync2(new(Student))
	if err !=nil{
		log.Fatal(err)
	}

}
