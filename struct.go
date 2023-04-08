package main

import "fmt"

type Point struct {
    X int
    Y int
}
type Person struct {
	hobby []string
	pos Point
}

func (p Point) Pointfunc() {
	fmt.Println(p.X , " " , p.Y)
}

func (p *Point) setpoint(x , y int) {
	p.X = x 
	p.Y = y
}

func main() {
	var p Point
	p.X = 10
	p.Y = 20	
	var p2 = &Point{}
	p2.X = 20
	p2.Y = 30
	var p3 = &Point{
		X: 30,
		Y: 30,
	}
	p3.setpoint(40 , 40)// 使用结构体方法
	p3.Pointfunc() 
	var son Person
	son.hobby = make([]string, 5 , 5)
	son.hobby[0] = "唱"
	son.hobby[1] = "跳"
	son.hobby[2] = "rap"
	son.pos.setpoint(5 , 5)
	
}
