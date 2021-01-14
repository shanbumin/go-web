package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		//把数字值推入通道c
		c <- i
		fmt.Println("Threw  >>", i)
	}
	return
}

//todo 上述仅仅写入5个就不写了，但是我们这边读多了也不会报错的，因为我们认为go thrower还会继续写的，哪怕for循环结束了
func catcher(c chan int) {
	for i := 0; i < 100; i++ {
		//从c中取出值
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(10000 * time.Millisecond)
}
