package main

import "fmt"

func producer(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("A produce", i)
		c <- i
	}
	defer close(c)
}
func consumer(c chan int) {
	hasMore := true
	//var p int		// 如果这里不定义p，就需要在后面使用 :=
	for hasMore { // 这样hasMore也被当成了局部变量,并且为bool型，每次for都会新建一个，默认值为true
		if p, hasMore := <-c; hasMore { // 一旦通道中没数据，就会阻塞，导致循环不能退出
			fmt.Printf("%v\n", hasMore)
			fmt.Println("B consume", p)
		}
	}
}

func main() {
	c := make(chan int)
	go producer(c)
	consumer(c)
}
