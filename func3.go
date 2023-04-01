package main

import (
    "fmt"
)

func fn() {
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println("err")
		}
	}()
	panic("抛出异常")
}

func div(a int , b int) int {
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println("err:",err)
		}
	}()
	return a / b
}

func main() {

    fmt.Println("defer begin")

    // 将defer放入延迟调用栈
    defer fmt.Println(1)

    defer fmt.Println(2)

    // 最后一个放入, 位于栈顶, 最先调用
    defer fmt.Println(3)

    fmt.Println("defer end")

	//异常
	fn()
	//除以0
	div(10 , 0)
}