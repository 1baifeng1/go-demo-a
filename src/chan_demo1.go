package main

import (
	"fmt"
	"time"
)

func main() {

}

// 使用channel实现超时机制
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
