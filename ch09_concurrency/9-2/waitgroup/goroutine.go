package main

import "fmt"
import "time"
import "sync"

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	//声明一个等待组
	var wg sync.WaitGroup
	//为计数器设置值
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	//阻塞到计数器的值为0
	wg.Wait()

	fmt.Println("终于结束了")
}
