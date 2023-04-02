package main

import (
    "fmt"
    "time"
)

func test() {
    start := time.Now() // 获取当前时间
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum++
    }
    elapsed := time.Since(start) //计算从某个时间点到当前时间的时间间隔
	// elapsed := time.Now().Sub(start) 计算两个时间点之间的时间间隔
    fmt.Println("该函数执行完成耗时：", elapsed)
}

func gettime () {
	start := time.Now() // 获取当前时间
	year := start.Year()
	month :=  start.Month()
	day := start.Day()
	hour := start.Hour()
	min := start.Minute()
	second :=  start.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,min,second)
	var str = start.Format("2006/01/02 03:04:05") //2006 1 2 3 3(15) 4 5
	fmt.Println(str)

	ti := start.Unix() //获取时间戳
	fmt.Println(ti)

	later := start.Add(time.Hour)//加一小时
    fmt.Println(later)
}

func main() {
    test()
	gettime()
}