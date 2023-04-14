package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test () {
	for i := 0; i < 3; i++ {
		fmt.Println("Test")
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器减一
}

func main() {
	wg.Add(1) //协程计数器加一
	go test() //开启协程 
	for i := 0; i < 3; i++ {
		fmt.Println("Main")
		time.Sleep(time.Millisecond)
	}
	wg.Wait()//主线进程完成后等待协程
}