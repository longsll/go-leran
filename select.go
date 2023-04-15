package main

import (
	"fmt"
)

func main() {
	intchannel := make(chan int , 10)
	stringchannel := make(chan string , 10)
	for i := 0; i < 10; i++ {
		intchannel <- i
	}
	for i := 0; i < 10; i++ {
		stringchannel <- "hello" + fmt.Sprintf("%d",i)
	}
	//select从intch与stringch同时读取数据
	for {
		select {
		case v := <- intchannel:
			fmt.Println("intchannel读取",v)
		case v := <- stringchannel:
			fmt.Println("stringchannel读取",v)
		default:
			fmt.Println("读取完毕")
			return
		}
	}
}