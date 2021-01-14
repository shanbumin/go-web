package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

//func main() {
//	a, b := make(chan string), make(chan string)
//	go callerA(a)
//	go callerB(b)
//	var msg string
//	openA, openB := true, true
//	for openA || openB {
//		select {
//		case msg, openA = <-a:
//			if openA {
//				fmt.Printf("%s from A\n", msg)
//			}
//		case msg, openB = <-b:
//			if openB {
//				fmt.Printf("%s from B\n", msg)
//			}
//		}
//	}
//}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	msg1, msg2 := "A", "B"
	for {
		//这里一定要间隔一下，否则协程还没有在管道中追加值就命中default了
		time.Sleep(1 * time.Microsecond)

		select {
		case msg1 = <-a:
			fmt.Printf("%s from A\n", msg1)
		case msg2 = <-b:
			fmt.Printf("%s from B\n", msg2)
		default:
			fmt.Println("Default")
		}
		if msg1 == "" && msg2 == "" {
			break
		}

	}
}
