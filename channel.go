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

func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("fn1" , i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch) //使用range要关闭channel
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Println("fn2",v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	var ch = make(chan int, 30)
	var _ = make(chan <- int, 5)//只写管道
	var _ = make(<- chan int, 5)//只读管道
	ch <- 100
	<- ch
	wg.Add(1) //协程计数器加一
	go test() //开启协程 
	for i := 0; i < 3; i++ {
		fmt.Println("Main")
		time.Sleep(time.Millisecond)
	}
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()//主线进程完成后等待协程
}