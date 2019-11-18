package main

import "fmt"

func main() {
	fmt.Println("(3)", fun1()) // 0
	fmt.Println("--------------------------")
	fmt.Println("(3)", fun2()) // 5
	fmt.Println("--------------------------")
	fmt.Println("(3)", fun3()) // 0

}

func fun1() int {
	i := 0
	fmt.Println("(1)", i)
	defer func() {
		i = 5
		fmt.Println("(2)", i)
	}()
	return i
}

func fun2() (i int) {
	i = 0
	fmt.Println("(1)", i)
	defer func() {
		i = 5
		fmt.Println("(2)", i)
	}()
	return i
}

func fun3() (i int) {
	i = 0
	fmt.Println("(1)", i)
	defer func(i int) { // 使用值传递，所以不会更改i的值
		i = 5
		fmt.Println("(2)", i)
	}(i)
	return
}
