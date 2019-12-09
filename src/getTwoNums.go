package main

import "fmt"

/*
# Question:
# Given an array of integers, return indices of the two numbers such that they add up to a specific target.
# You may assume that each input would have exactly one solution, and you may not use the same element twice.

# Example:
# Input: nums = [2, 7, 11, 15], target = 9
# Output: [0, 1]
# Because nums[0] + nums[1] = 2 + 7 = 9
*/
func main() {
	a := []int{2, 7, 11, 15}
	tar := 18
	fmt.Println(twoSum2(a, tar))

}

// Thinking: Use map structure in Go, we can store the nums into a map
func twoSum2(nums []int, target int) []int {
	num := make(map[int]int, len(nums))
	var a, b int
	for i := 0; i < len(nums); i++ {
		num[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]

		if j, ok := num[temp]; ok == true {
			a, b = i, j
			break
		}
	}
	return []int{a, b}
}
