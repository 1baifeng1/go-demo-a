package main

import "fmt"

func sliceAdd() {
	s := []int{1, 2, 3, 4, 5, 6}
	for i, e := range s {			// 切片是指针类型，e拿到改变后的数1，3，6...
		if i == len(s) - 1 {
			s[0] += e
		} else {
			s[i+1] += e
		}
	}

	fmt.Println(s)
	// Output: [22 3 6 10 15 21]
}
func arrayAdd() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	for i, e := range a {			// 数组是数值类型，e依次拿到1，2，3...
		if i == len(a) - 1 {
			a[0] += e
		} else {
			a[i+1] += e
		}
	}

	fmt.Println(a)
	// Output: [7 3 5 7 9 11]
}




func main() {
	sliceAdd()
	arrayAdd()
}
