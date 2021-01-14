package main

import "fmt"
import "time"

func printNumbers(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	//把一个布尔值放入通道，以便解除程序的阻塞状态
	w <- true
}

func printLetters(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	//把一个布尔值放入通道，以便解除程序的阻塞状态
	w <- true
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers(w1)
	go printLetters(w2)
	//主程序将一直阻塞，直到通道里面出现可弹出的值为止
	<-w1
	<-w2
}
