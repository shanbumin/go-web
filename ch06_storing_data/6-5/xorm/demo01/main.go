package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)
/*
CREATE TABLE `user` (
  `id`    bigint(20) NOT NULL AUTO_INCREMENT,
  `name`  varchar(255) DEFAULT NULL,
  `salt`  varchar(255) DEFAULT NULL,
  `age`   int(11) DEFAULT NULL,
  `passwd` varchar(200) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

 */

type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

/*
CREATE TABLE `detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_detail_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type Detail struct {
	Id int64
	UserId int64 `xorm:"index"`
}


func main() {

	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	//第一步创建引擎
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
      log.Fatal(err)
	}

    //第二步 定义一个和表同步的结构体，并且自动同步结构体到数据库
	err = engine.Sync2(new(User))

	if err !=nil{
		log.Fatal(err)
	}

	err=engine.Sync2(new(Detail))
	if err !=nil{
		log.Fatal(err)
	}


}
