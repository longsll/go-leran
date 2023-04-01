package main

import "fmt"

func myfunc(args ...int) int {
	res := 0
    for _, arg := range args {
        res += arg
    }
	return res
}

func main() {
	//arr  := []int { 1 , 1 , 2 ,5}
	x := myfunc(5 , 6 ,7)
	fmt.Println(x)
}