package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


//对文件进行读写

func main() {

	data := []byte("Hello World!\n")
	// ====================write to file and read from file using WriteFile and ReadFile
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))
	// =====================除了ioutil库，还可以使用os库的函数进行文件读写操作。
	//写
	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	//fmt.Printf("Wrote %d bytes to file\n", bytes)
    //读
	file2, _ := os.Open("data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))


}
