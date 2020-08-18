package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

//联合普通索引
//唯一索引
//普通索引
//默认值
//非空
//...



//使用 xorm:"'column_name'"可以使该 field 对应的 Column 名称为指定名称。这里使用两个单引号将 Column 名称括起来是为了防止名称冲突，
//因为我们在Tag中还可以对这个Column进行更多的定义。如果名称不冲突的情况，单引号也可以不使用。

type Student struct {
	Id        int64 `xorm:"pk autoincr"`
	Age      int     `xorm:"notnull default 0 comment('年纪')"`
	Name     string  `xorm:"varchar(255) notnull unique  default '' comment('姓名')"`
	Url       string  `xorm:"index notnull index default '' comment('地址')"`
	Status    uint8    `xorm:"notnull default 0 comment('状态')"`
	Version   int64   `xorm:"version"`
	SpecialCreatedAt   int64 `xorm:"created"`
	CreatedAt time.Time `xorm:"created"` //todo 注意不要用单引号括起来created,单引号是用来括字段名，当字段名与行为符冲突的时候额
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt time.Time `xorm:"'deleted'"` // 此特性会激发软删除
}

func (m *Student) TableName() string {
	return "student3"
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

	//插入
	stu01:=&Student{
		Name:"sam",
		Age:11,
		Status:1,
	}
	engine.Insert(stu01)

}
