package main

import "fmt"

func main() {
	// create methords
	//var s1 []int					// nil slice: array=0, len=0, cap=0
	//s12 := *new([]int)			// nil slice: same to above
	//var s2 = []int{}				// empty slice: array=0x42003bda0, len=0, cap=0
	//var s22 = make([]int, 0)		// empty slice: same to above
	s3 := []int{1, 2, 3, 4, 5} // array=0x..., len=5, cap=5
	fmt.Println("s3: len=", len(s3), "cap=", cap(s3))
	s4 := make([]int, 5, 8) // array=0x..., len=5, cap=8
	fmt.Printf("s4: len=%d, cap=%d \n", len(s4), cap(s4))
	s5 := s3[2:4] // array=[3, 4] len=2, cap=3	// 前闭后开，包前不包后
	fmt.Println("s5=", s5, ": len=", len(s5), " cap=", cap(s5))

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]
	fmt.Println("s1")

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)

	ModifyANum(s1)
	fmt.Println(s1)
}

func ModifyANum(s []int) { // 切片作为参数传递时，引用传递
	s[0] = 123321
}
