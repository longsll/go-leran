package main

import "fmt"

func change1(s byte) int { //字符'1'转数字1
	return int(s) - 48
}

func main() {
	fmt.Print(change1('1'))
}
