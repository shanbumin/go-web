package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)


//名称映射规则主要负责结构体名称到表名和结构体 field 到表字段的名称映射。


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

/*
CREATE TABLE `student` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`age` int(11) DEFAULT NULL,
`name` varchar(255) DEFAULT NULL,
`img_url` varchar(255) DEFAULT NULL,
`url` varchar(255) DEFAULT NULL,
`status` int(11) DEFAULT NULL,
`created_at` datetime DEFAULT NULL,
`updated_at` datetime DEFAULT NULL,
`deleted_at` datetime DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


insert into student (age,name,img_url,url,status,created_at,updated_at)values
(38,"sam01","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(28,"sam02","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(28,"sam03","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(28,"sam04","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(38,"sam05","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(28,"sam06","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(38,"sam07","https://www.baidu.com/a.png","www.baidu.com",1,now(),now()),
(28,"sam08","https://www.baidu.com/a.png","www.baidu.com",1,now(),now());



*/
type Student struct {
	Id        int64 `xorm:"pk autoincr"` // 注：使用getOne 或者ID() 需要设置主键
	Age      int
	Name     string
	ImageUrl  string `xorm:"'img_url'"`  //使用 xorm:"'column_name'"可以使该 field 对应的 Column 名称为指定名称
	Url       string
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `xorm:"deleted"` // 此特性会激发软删除
}
//todo 当表名需要特殊指明的时候可以在这里，比如与结构体不一样的时候students
//todo  这里指明了表名将覆盖引擎统一设置的表前缀额
//todo 如果结构体拥有 TableName() string 的成员方法，那么此方法的返回值即是该结构体对应的数据库表名。
func (m *Student) TableName() string {
	return "student"
}

func main() {

	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	//第一步创建引擎
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
      log.Fatal(err)
	}

	err=engine.Ping()
	//err=engine.PingContext()
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("connect success.")
	}

	//加个统一的表前缀试试
	//names.NewPrefixMapper
	//names.NewSuffixMapper 表后缀
	//names.NewCacheMapper
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "s_")
	engine.SetTableMapper(tbMapper)



	//todo  engine 可以通过 engine.Close 来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。
	defer engine.Close()
    //第二步 定义一个和表同步的结构体，并且自动同步结构体到数据库
	err = engine.Sync2(new(User))
	if err !=nil{
		log.Fatal(err)
	}

	err=engine.Sync2(new(Detail))
	if err !=nil{
		log.Fatal(err)
	}


	err=engine.Sync2(new(Student))
	if err !=nil{
		log.Fatal(err)
	}




}
