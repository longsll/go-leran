package main

import "fmt"

type usb interface { //用接口实现规范
	start()
	stop()
}

type iphone struct {
	Name string
}

func (p iphone)start() {
	fmt.Println("开机")
}

func (p iphone)stop() {
	fmt.Println("关机")
}

type computer struct {

}

func (c computer) wokr(usber usb) {
	usber.start()
	usber.stop()
}

func main () {

	p := iphone {
		Name: "huawei",
	}

	var p1 usb
	p1 = p

	p1.start()

	//computer调用iphone的usb接口
	var comp = computer{}
	comp.wokr(p1)

	//空接口可以表示任何类型
	var a interface{}
	a = 20
	fmt.Printf("%v , %T\n" , a , a)
	a = "rap"
	fmt.Printf("%v , %T\n" , a , a)
	v , ok := a.(string) //类型断言
		if ok { //判断成功
			fmt.Println(v)
		}
}