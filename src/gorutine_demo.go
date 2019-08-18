package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, elem := range a {
		sum += elem
	}
	c <- sum
}

func test_sum() {
	array := []int{4, 6, -9, 2, 1, 0, -4}

	c := make(chan int)
	go sum(array[:len(array) / 2], c)
	go sum(array[len(array) / 2 : ], c)

	//r1, r2 := <-c, <-c
	//fmt.Println(r1, r2, r1 + r2)
	r1 := <-c
	r2 := <-c
	fmt.Println(r1)
	fmt.Println(r2)
}
/***********************************/

type worker struct{
	channel chan int
	done chan bool
}

func test_worker(){
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = worker{
			make(chan int, 1),
			make(chan bool),
		}
		go printWorker(i, workers[i])
	}

	for i, w := range workers {
		w.channel <- 10 + i
	}

	//for _, w := range workers {
	//	<- w.done
	//}

	for i, w := range workers {
		w.channel <- 40 + i
	}

	for _, w := range workers {
		<- w.done
		<- w.done
	}
}

func printWorker(i int, w worker){
	for {
		fmt.Println(i, "got", <-w.channel)
		w.done <- true
		//go func() {w.done<-true}()
	}
}

///////////////////////////////////////////////////
func main() {
	//test_sum()
	test_worker()
}
