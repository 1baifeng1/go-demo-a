package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (stack *Stack) Push(elem int) {
	*stack = append(*stack, elem)
}

func (stack *Stack) Pop() int {
	var top int
	length := len(*stack)
	if length > 0 {
		st := *stack
		top = st[length-1]
		if length > 1 {
			*stack = append(st[:length-1])
		} else {
			*stack = nil
		}
	}
	return top
}

func (stack *Stack) GetMin() int {
	min := 9999999
	st := *stack
	for i := 0; i < len(st); i++ {
		if st[i] < min {
			min = st[i]
		}
	}
	return min
}

type Queue struct {
	s1 Stack
	s2 Stack
}

func (queue *Queue) Add(x int) {
	queue.s1.Push(x)
}

func (queue *Queue) Poll() {
	for len(queue.s1) > 1 {
		queue.s2.Push(queue.s1.Pop())
	}
	queue.s1.Pop()
	for len(queue.s2) > 0 {
		queue.s1.Push(queue.s2.Pop())
	}
}

func (queue *Queue) Peek() int {
	for len(queue.s1) > 1 {
		queue.s2.Push(queue.s1.Pop())
	}
	x := queue.s1.Pop()
	queue.s1.Push(x)
	for len(queue.s2) > 0 {
		queue.s1.Push(queue.s2.Pop())
	}

	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _, err := reader.ReadLine()
	if err != nil {
		return
	}
	lines, _ := strconv.Atoi(string(text))

	// 设计getMin功能的栈
	//var stack Stack
	//for i := 0; i < lines; i++ {
	//	text, _, _ = reader.ReadLine()
	//	t := string(text)
	//	if strings.HasPrefix(t, "push") {
	//		elems := strings.Fields(t)
	//		e, err := strconv.Atoi(elems[1])
	//		if err != nil {
	//			return
	//		}
	//		stack.Push(e)
	//	} else if strings.HasPrefix(t, "pop") {
	//		//fmt.Println(stack.Pop())
	//		stack.Pop()
	//	} else if strings.HasPrefix(t, "getMin") {
	//		fmt.Println(stack.GetMin())
	//	}
	//}

	// 由两个栈组成的队列
	var queue Queue
	for j := 0; j < lines; j++ {
		text, _, _ := reader.ReadLine()
		t := string(text)

		if strings.HasPrefix(t, "add") {
			elems := strings.Fields(t)
			e, err := strconv.Atoi(elems[1])
			if err != nil {
				return
			}
			queue.Add(e)
		} else if strings.HasPrefix(t, "poll") {
			queue.Poll()
		} else if strings.HasPrefix(t, "peek") {
			fmt.Println(queue.Peek())
		}
	}
}
