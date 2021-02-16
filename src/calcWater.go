package main

import "fmt"

func main() {
	sum := calcWater([]int{1, 1, 5, 2, 3, 2, 5})
	fmt.Println(sum)
}

func calcWater(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return 0
	}

	sum := 0
	for i := 1; i < len(arr) - 1; i++ {
		if arr[i - 1] > arr[i] {
			for j := i; j < len(arr) - 1; j++ {

			}
		}

		if arr[i - 1] > arr[i] && arr[i + 1] > arr[i] {
			min := 0
			if arr[i - 1] <= arr[i + 1]{
				min = arr[i - 1] - arr[i]
			} else {
				min = arr[i + 1] - arr[i]
			}
			sum += min
		}
	}


	return sum
}
