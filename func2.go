package main

import "fmt"

type cacl func(int , int) int

func add(x , y int) int {
	return x + y
}

//闭包
// Go语言中的闭包是指一个函数和它所引用的外部变量组合成的一个整体。闭包可以访问外部函数中的变量
// 并且这些变量的值在闭包中被保留
// 即使外部函数已经返回，闭包仍然可以访问这些变量
// 这里定义了一个`outer`函数，它返回一个匿名函数。
// 匿名函数中访问了`outer`函数中的变量`i`，并且将`i`的值加1后返回。
// 由于匿名函数中引用了`outer`函数中的变量，因此匿名函数和变量`i`组成了一个闭包
func outer() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

	//函数作为参数
	var c cacl
	c = add
	fmt.Println(c(2 , 5))

	//匿名函数,可在另一个函数里使用
	func(data int) {
		fmt.Println("hello", data)
	}(100)
	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	f(100)

	// 这里首先调用`outer`函数，得到一个匿名函数`g`，然后多次调用`g`函数，每次调用都会将变量`i`的值加1并返回。
	// 由于变量`i`的值在闭包中被保留，因此每次调用`g`函数时都会使用上一次调用时的变量`i`的值
	g := outer()
    fmt.Println(g()) // 输出1
    fmt.Println(g()) // 输出2
    fmt.Println(g()) // 输出3
}