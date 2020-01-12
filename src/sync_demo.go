package main

import (
	"fmt"
	"sync"
	"time"
)

// Sync包的使用
/*
	1. sync.Once() 对指定函数，多次调用，仅单次执行
	2. sync.WaitGroup{} 组等待
	3. sync.Mutex{} and sync.RWMutex{} 互斥锁和读写互斥锁(互斥锁是忙睡眠的)
	4. sync.Cond 条件变量
*/

func main() {
	//onceDemo()
	//waitGroupDemo()
	//mutexDemo()
	//rwmutextDemo()
	condDemo()
}

func onceDemo() {
	var once1 sync.Once
	//var once2 sync.Once
	once1Func := func() {
		fmt.Println("aaaaaaaaaaaaaa")
	}
	//once2Func := func(i int) {
	//	fmt.Println("bbbbbbbbbbbbbb", i)
	//}

	flag := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once1.Do(once1Func)
			flag <- true

			//if i > 5 {
			//	once2.Do(once2Func(i))  // 不支持带参数的回调函数
			//	//once2Func(i)			// 由于是在协程中执行，传入的i在使用时可能已经成10了
			//}
		}()
	}
	for i := 0; i < 10; i++ {
		<-flag
	}
}

func waitGroupDemo() {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1) // 可以在新建协程时添加
		go func() {
			fmt.Println(i)
			waitGroup.Done() // 每个协程结束可以设置为done
		}()
	}
	waitGroup.Wait()
	fmt.Println("All routines run finished.")
}

type ClockInt struct {
	sync.Mutex
	data int
}

func mutexDemo() {
	myClockInt := ClockInt{}
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) { // 将i传入，可以避免取到错误的值
			myClockInt.Lock()
			myClockInt.data += i
			fmt.Println(myClockInt.data)
			myClockInt.Unlock()
			ch <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch // 使用一个通道阻塞等待也可以实现主协程等待其它协程退出后退出
	}
}

type ClockInt2 struct {
	sync.RWMutex
	data int
}

func rwmutextDemo() {
	myClockInt := ClockInt2{}
	myClockInt.data = 1
	wait := sync.WaitGroup{}
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			if j%2 == 0 {
				myClockInt.Lock() // 加写锁	// 同一时间只能有一个写锁
				myClockInt.data++
				fmt.Println("Write lock")
				fmt.Println(myClockInt.data)
				time.Sleep(1)
				fmt.Println("Write unlock")
				myClockInt.Unlock() // 解写锁
			} else {
				myClockInt.RLock() // 加读锁	// 可以同时有多个读锁
				fmt.Println("Read lock")
				fmt.Println(myClockInt.data)
				time.Sleep(1)
				fmt.Println("Read unlock")
				myClockInt.RUnlock() // 解读锁
			}
			wait.Done()
		}(i)
	}
	wait.Wait()
}

func condDemo() {
	var queue int8 // 只包含1个值的，0和1代表长度
	mutex := sync.RWMutex{}
	writeCond := sync.NewCond(&mutex) // 两个条件，有相同的临界区
	readCond := sync.NewCond(&mutex)
	write := func(index int) { // 生产者
		mutex.Lock()
		for queue == 1 { // 如果队列满，就阻塞等待
			writeCond.Wait()
		}
		fmt.Println("Index", index, "write message into queue.")
		queue = 1
		readCond.Broadcast() // 有多个消费者，需要广播
		mutex.Unlock()
	}
	read := func(index int) { // 消费者
		mutex.Lock()
		for queue == 0 { // 如果队列空，就阻塞等待
			readCond.Wait() // Wait()会进行Unlock()操作，所以之前必须加锁，而且不能是加读锁
		}
		fmt.Println("Index", index, "read message from queue.")
		queue = 0
		writeCond.Signal() // 有一个生产者，只需发信号
		mutex.Unlock()
	}

	// 创建一个生产者，两个消费者
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	go func() {
		for i := 0; i < 6; i++ {
			write(0)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			read(0)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			read(1)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
