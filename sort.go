package main

import (
	"fmt"
	"sort"
)


type Person struct {
	Name string    // 姓名
	Age  int    // 年纪
	Price int // 学费
}
 
// type PersonSlice [] Person
// func (a PersonSlice) Len() int {    // 重写 Len() 方法
// 	return len(a)
// }
// func (a PersonSlice) Swap(i, j int){     // 重写 Swap() 方法
// 	a[i], a[j] = a[j], a[i]
// }
// func (a PersonSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
// 	return a[i].Price < a[j].Price
// }


func main() {

	people := [] Person{
		{"zhang san", 12, 20},
		{"li si", 30, 20},
		{"wang wu", 12,18},
		{"zhao liu", 26, 15},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按照年龄从低到高排序的结果\n",people)


	a := []int{6, 3, 9, 8, 1, 2, 5, 7}
	sort.Ints(a) //sort.Float64s()  //sort.Strings() 
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}

