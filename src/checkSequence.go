package main

import "fmt"

/*
[
Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.
Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.
]
 */


func checkSequence(array []int) bool {
	var num int = 0
	var flag bool

	if len(array) <= 2 {
		return true
	}
	for i := 0; i < len(array)-1; i++ {
		if array[i] >= array[i + 1]{
			num++
			//fmt.Println(i)
		}
	}

	if num > 1 {
		flag = false
	} else {
		flag = true
	}

	return flag
}

func main() {
	array := []int {1, 10, 3, 5, 7, 9}		// T
	array1 := []int {1}						// T
	array2 := []int {2, 1}					// T
	array3 := []int {1, 3, 5, 4, 7, 9}		// T
	array4 := []int {1, 3, 5, 7, 9}			// T
	array5 := []int {1, 3, 5, 7, 6, 7, 9}	// F
	fmt.Println(checkSequence(array))
	fmt.Println(checkSequence(array1))
	fmt.Println(checkSequence(array2))
	fmt.Println(checkSequence(array3))
	fmt.Println(checkSequence(array4))
	fmt.Println(checkSequence(array5))
}
