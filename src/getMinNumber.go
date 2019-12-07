package main

import (
	"fmt"
)

// 给定一个大整数，从其中去掉n个数字，使剩下的数字组成的整数最小

// 思路：
// 1、从最高位开始遍历，当前位和下一位比较，如果当前位的数字大则去掉当前位；
//    如果下一位数字大则将下一位变为当前位，如果没有下一位，则删除当前位
//    遍历n次
func GetRemainMaxNumSli(ori []int, n int) int {
	defer func() int {
		if errNum := recover(); errNum != nil {
			if errNum == 1 {
				fmt.Println("invalid input")
			} else if errNum == 2 {
				fmt.Println("request a big integer")
			}
		}
		return 0
	}()

	if n >= len(ori) {
		panic(1)
	} else if len(ori) == 1 {
		panic(2)
	}

	for j := 0; j < n; j++ {
		preNum := ori[0]
		for i := 1; i < len(ori); i++ {
			curNum := ori[i]
			if preNum > curNum {
				ori = RemoveAFromSlice(ori, i-1)
				break
			} else if i == len(ori)-1 {
				// If it is the last one, then it's the max number
				ori = RemoveAFromSlice(ori, i)
			} else {
				preNum = curNum
			}
		}
	}

	result := 0
	for i := 0; i < len(ori); i++ {
		result = result*10 + ori[i]
	}
	return result
}

func RemoveAFromSlice(sli []int, index int) []int {
	// slice is value passing, we can use pointer to pass
	if index == 0 {
		sli = sli[1:]
	} else if index == len(sli)-1 {
		sli = sli[:index]
	} else {
		sli = append(sli[:index], sli[index+1:]...)
	}

	return sli
}

func SplitIntToSlice(num int) (sli []int) {
	// High digit needs to be set to low Slice
	rest := num / 10
	rema := num % 10
	//fmt.Println(rest, rema)
	for rest != 0 {
		te := []int{rema}
		sli = append(te, sli...)
		rema = rest % 10 // Take the rest first
		rest = rest / 10 // Then divide
		//fmt.Println(rest, rema)
	}
	sli = append([]int{rema}, sli...)
	fmt.Println(sli)
	return
}

// 2、记录一个数作为最小值，每次去掉一个数，剩下的数和最小值比较，如果剩下的数较小就替换最小值，否则继续遍历

func main() {
	test1 := 1234567
	test2 := 21345
	test3 := 2025433
	test4 := 123454321
	test5 := 122232

	fmt.Printf("%d remove %d int, remain is %d", test1, 2, GetRemainMaxNumSli(SplitIntToSlice(test1), 2))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test2, 1, GetRemainMaxNumSli(SplitIntToSlice(test2), 1))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test3, 1, GetRemainMaxNumSli(SplitIntToSlice(test3), 1))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test4, 2, GetRemainMaxNumSli(SplitIntToSlice(test4), 2))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test5, 1, GetRemainMaxNumSli(SplitIntToSlice(test5), 1))
	fmt.Println()
	fmt.Println("===================================================================")
	fmt.Printf("%d remove %d int, remain is %d", test1, 2, GetRemainMaxNumPointer(SplitIntToPointer(test1), 2))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test2, 1, GetRemainMaxNumPointer(SplitIntToPointer(test2), 1))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test3, 1, GetRemainMaxNumPointer(SplitIntToPointer(test3), 1))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test4, 2, GetRemainMaxNumPointer(SplitIntToPointer(test4), 2))
	fmt.Println()
	fmt.Printf("%d remove %d int, remain is %d", test5, 1, GetRemainMaxNumPointer(SplitIntToPointer(test5), 1))

}

// 将数字拆成指针
type IntNode struct {
	Data int
	Next *IntNode
}

func SplitIntToPointer(num int) *IntNode {
	rest := num / 10
	remainder := num % 10
	head := &IntNode{remainder, nil}
	for rest != 0 {
		remainder = rest % 10
		rest = rest / 10
		node := &IntNode{remainder, nil}
		node.Next = head
		head = node
	}
	for h := head; h != nil; h = h.Next {
		fmt.Print(h.Data)
	}
	fmt.Println()
	return head
}

func GetRemainMaxNumPointer(head *IntNode, n int) (result int) {
	if head == nil {
		return -1
	}

	for j := 0; j < n; j++ {
		var pre *IntNode = nil
		flag := false // Record if remove a number
		for h := head; h.Next != nil; h = h.Next {
			if h.Data > h.Next.Data {
				if pre == nil {
					head = h.Next
				} else {
					pre.Next = h.Next
				}
				flag = true
				break
			} else {
				pre = h
			}
		}
		if !flag {
			pre.Next = nil
		}

	}

	for h := head; h != nil; h = h.Next {
		result = result*10 + h.Data
	}

	return
}
