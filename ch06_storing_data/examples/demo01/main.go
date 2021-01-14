package main
import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
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

func ping(){
	err := db.Ping()
	if err != nil{
		log.Fatalln(err)
	}
}

func  createTableHello(){
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil{
		log.Fatalln(err)
	}

}

func  insertTableHello(){
	rs, err := db.Exec("INSERT INTO test.hello(world) VALUES ('hello world')")
	if err != nil{
		log.Fatalln(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil{
		log.Fatalln(err)
	}
	log.Printf("inserted %d rows", rowCount)
}


//我们使用了Query方法执行select查询语句，返回的是一个sql.Rows类型的结果集。
//迭代后者的Next方法，然后使用Scan方法给变量s赋值，以便取出结果。最后再把结果集关闭（释放连接）。
func queryTableHello(){
	rows, err := db.Query("SELECT world FROM test.hello")
	if err != nil{
		log.Fatalln(err)
	}

	for rows.Next(){
		var s string
		err = rows.Scan(&s)
		if err !=nil{
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
}


func main() {

	     defer db.Close()
	     //ping
	     ping()
		//创建一个数据表
		createTableHello()
	    //此时可以看见，数据库生成了一个新的表。接下来再插入一些数据。
	    //同样使用Exec方法即可插入数据，返回的结果集对象是是一个sql.Result类型，它有一个LastInsertId方法，返回插入数据后的id。
	    //当然此例的数据表并没有id字段，就返回一个0.
        insertTableHello()
		//查询
		queryTableHello()






}

