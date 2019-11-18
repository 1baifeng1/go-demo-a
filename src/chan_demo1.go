package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

}

// 使用channel和select实现超时机制
func timeout(ch chan interface{}) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-timeout:
		fmt.Println("It's timeout.")
		return
	}
}

// 使用select的default检测channel buffer是否已满
func isBufferFull() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("ch")
	default:
		fmt.Println("ch is full.")
	}
}
func IsBufferFull(ch chan interface{}, x interface{}) {
	select {
	case ch <- x:
		fmt.Printf("%v\n", ch)
	default:
		fmt.Println("It's full.")
	}
}

// 使用sync.WaitGroup监控子协程的执行完成情况
func MonitorChannel() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("goroutine2")
		wg.Done()
	}()
	wg.Wait()
}
