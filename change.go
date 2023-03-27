package main

import "fmt"

func change1(s byte) int { //字符'1'转数字1
	return int(s) - 48
}

func change2(s byte) int { //字符'a'转数字1
	return int(s) - 96
}

func change3(t int) byte { //数字1转字符'1'
	return byte(t + 48)
	//fmt.Printf("%c",change3(1))
}

func change4(t int) byte { //数字1转字符'a'
	return byte(t + 96)
	//fmt.Printf("%c",change4(1))
}

func main() {
	fmt.Printf("%c",change4(1))
}
