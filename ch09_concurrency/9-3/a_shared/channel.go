package main

import (
	"fmt"
	// "math/rand"
	"time"
	"runtime"
)

var DB Store

type Store struct {
	hash map[string]string
	in chan [2]string //通道元素是简单的有长度为2的数组
	out chan [2]string
}

func StoreInit() {
	DB = Store{
		hash: make(map[string]string),
		in: make(chan [2]string),
	}
	go func() {
		for {
			a := <-DB.in
			DB.hash[a[0]] = a[1]
		}
	}()
}
func (store *Store) Get(key string) (value string, err error) {
	value = store.hash[key]
	return
}
func (store *Store) Add(key string, value string) (err error) {
	a := [2]string{key, value}
	store.in <- a
	// store.hash[key] = value
	return
}
//----------------------------------------------------------------------
func main() {
	runtime.GOMAXPROCS(4)
	StoreInit()
	for i := 0; i < 10; i++ {
		go DB.Add("a", "A")
		go DB.Add("a", "B")
		go DB.Add("a", "C")
		time.Sleep(1 * time.Microsecond)
		s, _ := DB.Get("a")
		fmt.Printf("%s\r\n ", s)
		
	}
}
