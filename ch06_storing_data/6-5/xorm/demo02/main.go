package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)


//创建Engine组
//@todo 所有使用 engine 都可以简单的用 engineGroup 来替换。
func main() {

	driverName:="mysql"
	masterDataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	slave1DataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	slave2DataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"


	//创建Engine组
	dataSourceNameSlice := []string{masterDataSourceName, slave1DataSourceName, slave2DataSourceName}
	engineGroup, err := xorm.NewEngineGroup(driverName, dataSourceNameSlice)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(engineGroup)
	//创建Engine组(二)
	//masterEngine, err := xorm.NewEngine(driverName, masterDataSourceName)
	//slave1Engine, err := xorm.NewEngine(driverName, slave1DataSourceName)
	//slave2Engine, err := xorm.NewEngine(driverName, slave2DataSourceName)
	//engineGroup, err := xorm.NewEngineGroup(masterEngine, []*Engine{slave1Engine, slave2Engine})
}
