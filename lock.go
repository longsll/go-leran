package main

import (
	"fmt"
	"sync"
	"time"
)


var  wg sync.WaitGroup

type Counter struct {
    count int
    mutex sync.Mutex
}

func (c *Counter) Increment() { //用于增加
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.count++
	fmt.Println("add",c.count)
	wg.Done()
}

func (c *Counter) Decrement() { //用于减少
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.count--
	time.Sleep(time.Millisecond * 50)
	fmt.Println("cmp",c.count)
	wg.Done()
}

func (c *Counter) Value() int { //输出值
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.count
}

func main() {
    counter := Counter{}

	//同时开启10个增加和10个减少进程
    for i := 0; i < 10; i++ {
		wg.Add(1)
        go counter.Increment()
    }
    for i := 0; i < 10; i++ {
		wg.Add(1)
        go counter.Decrement()
    }
	
	wg.Wait()
	fmt.Println(counter.Value())
}