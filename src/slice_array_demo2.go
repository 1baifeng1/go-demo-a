package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	fmt.Println("s1:", len(s1), cap(s1))	// s1: 3 8
	fmt.Println("s2:", len(s2), cap(s2))	// s2: 4 5

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)			// [2 3 20]
	fmt.Println(s2)			// [4 5 6 7 100 200]
	fmt.Println(slice)		// [0 1 2 3 20 5 6 7 100 9]
}
