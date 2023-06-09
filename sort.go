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

func sortMapByKey(m map[int]int) {
    keys := make([]int, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        fmt.Printf("%d:%d\n", k, m[k])
    }
}

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

	//mapsort
	scene := make(map[int]int)
	scene[6] = 66
	scene[3] = 4
	scene[9] = 960
	sortMapByKey(scene)

	//该函数使用二分查找的方法，会从[0, n)中取出一个值index，index为[0, n)中最小的使函数f(index)为True的值
	//并且f(index+1)也为True。如果无法找到该index值，则该方法为返回n。
	aa := []int{1,2,3,4,5}
    d := sort.Search(len(aa), func(i int) bool { return aa[i]>3 })
    fmt.Println(d)
}
