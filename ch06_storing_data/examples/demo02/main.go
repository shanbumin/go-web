package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var db *sql.DB

func init(){
	var err error
	//创建一个数据库抽象对象
	db, err= sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil{
		log.Fatal(err)
	}

}



//读取数据
func query001(){
	//执行查询的sql语句,并将数据库连接的所属权传递给Result结果集
	rows, err := db.Query("SELECT world FROM test.hello")
	if err != nil{
		log.Fatalln(err)
	}
	defer rows.Close()



	//1.rows.Next方法设计用来迭代。当它迭代到最后一行数据之后，会触发一个io.EOF的信号，即引发一个"错误"，同时go会自动调用rows.Close方法释放连接，然后返回false。此时循环将会结束退出。
	//2.通常你会正常迭代完数据然后退出循环。可是如果并没有正常的循环而因其他错误导致退出了循环。此时rows.Next处理结果集的过程并没有完成，归属于rows的连接不会被释放回到连接池。
	//因此十分有必要正确的处理rows.Close事件。如果没有关闭rows连接，将导致大量的连接并且不会被其他函数重用，就像溢出了一样。最终将导致数据库无法使用。


	for rows.Next(){
		var s string
		err = rows.Scan(&s)
		if err !=nil{
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
	//rows.Next循环迭代的时候，因为触发了io.EOF而退出循环。为了检查是否是迭代正常退出还是异常退出，需要检查rows.Err。例如上面的代码应该改成：
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}


//读取单条记录
//Query方法是读取多行结果集，实际开发中，很多查询只需要单条记录，不需要再通过Next迭代。golang提供了QueryRow方法用于查询单条记录的结果集。
func  query002(){
	var s string
	var err error

	//QueryRow方法的使用很简单，它要么返回sql.Row类型，要么返回一个error，如果是发送了错误，则会延迟到Scan调用结束后返回，如果没有错误，则Scan正常执行。
	//只有当查询的结果为空的时候，会触发一个sql.ErrNoRows错误。你可以选择先检查错误再调用Scan方法，或者先调用Scan再检查错误。

	err = db.QueryRow("SELECT world FROM test.hello LIMIT 1").Scan(&s)//结果集方法Scan可以把数据库取出的字段值赋值给指定的数据结构。它的参数是一个空接口的切片，这就意味着可以传入任何值。
	                                                                        // 通常把需要赋值的目标变量的指针当成参数传入，它能将数据库取出的值赋值到指针值对象上。
	if err != nil{
		if err == sql.ErrNoRows{
			log.Println("There is not row")
		}else {
			log.Fatalln(err)
		}
	}
	log.Println("found a row", s)
}

//空值处理

/*
mysql> select * from hello;
+-------------+----+
| world       | id |
+-------------+----+
| hello world |  3 |
| hello world |  0 |
| hello world |  0 |
| NULL        |  4 |
+-------------+----+
4 rows in set (0.01 sec)
*/

func  query003(){

	//数据库有一个特殊的类型，NULL空值。可是NULL不能通过scan直接跟普遍变量赋值，甚至也不能将null赋值给nil。对于null必须指定特殊的类型，这些类型定义在database/sql库中。
	//例如sql.NullFloat64。如果在标准库中找不到匹配的类型，可以尝试在驱动中寻找。

	var world sql.NullString
	err:= db.QueryRow("SELECT world FROM hello WHERE id = ?",4).Scan(&world)

	fmt.Println(err)
	if world.Valid {
		fmt.Println("不是null",world.String)
	} else {
		// 数据库的value是不是null的时候，输出 world的字符串值， 空字符串
		fmt.Println("null",world.String)
	}

	//对应的，如果world字段是一个int，那么声明的目标变量类似是sql.NullInt64，读取其值的方法为world.Int64。

}

//但是有时候我们并不关心值是不是Null,我们只需要吧他当一个空字符串来对待就行。这时候我们可以使用[]byte（null byte[]可以转化为空string）
func  query004(){

	var world []byte
	err := db.QueryRow("SELECT world FROM hello WHERE id = ?",3).Scan(&world)
	fmt.Println(err)
	fmt.Println(string(world)) // 有值则取出字串值，null则转换成 空字串。
}


//在执行查询的时候，我们定义了目标变量，同时查询的时候也写明了字段，如果不指名字段，或者字段的顺序和查询的不一样，都有可能出错。
//因此如果能够自动匹配查询的字段值，将会十分节省代码，同时也易于维护。
//go提供了Columns方法用获取字段名，与大多数函数一样，读取失败将会返回一个err，因此需要检查错误。

/*
mysql> select * from user;
+----+------+-----------+
| id | gid  | real_name |
+----+------+-----------+
|  1 | 1000 | sam       |
|  2 | NULL | tom       |
|  3 | 1001 | NULL      |
|  4 | 1002 | good      |
|  5 | NULL | NULL      |
+----+------+-----------+
5 rows in set (0.00 sec)

 */

func  query005(){

	rows, err := db.Query("SELECT * FROM user WHERE gid = 1000")
	if err != nil{
		log.Fatalln(err)
	}
	defer rows.Close()


	cols, err := rows.Columns()
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(cols)//[id gid real_name]

	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range vals{
		scans[i] = &vals[i]
	}

	var results []map[string]string

	for rows.Next(){
		err = rows.Scan(scans...)
		if err != nil{
			log.Fatalln(err)
		}

		row := make(map[string]string)
		for k, v := range vals{
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}

	for k, v :=range results{
		fmt.Println(k, v)
	}




}




func main() {

	defer db.Close()

    //query001()
    //query002()
    //query003()
    //query004()
    query005()







}



