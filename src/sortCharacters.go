package main

/*
https://www.nowcoder.com/practice/d9aa3894d3aa4887843a85d26daa4437?tpId=61&&tqId=29548&rp=1&ru=/activity/oj&qru=/ta/pku-kaoyan/question-ranking
 */

import (
	"bufio"
	"fmt"
	"os"
)

var(
	n_sortChara int
)

func main() {
	fmt.Scanf("%d", &n_sortChara)
	inputReader := bufio.NewReader(os.Stdin)
	for i := 0; i < n_sortChara; i++ {
		var chara string
		//fmt.Scanln(&chara)
		chara, err := inputReader.ReadString('\n')
		if err == nil {
			//output := sortCharacters(chara)
			//fmt.Println(output)

			output1 := sortCharaBucket(chara)
			fmt.Println(output1)
		}
	}
}

func sortCharacters(chara string) (string) {
	var index int
	var output string
	var indexArray []int

	for i := 0; i < len(chara); i++ {
		if (chara[i] >= 97 && chara[i] <= 122) || (chara[i] >= 65 && chara[i] <= 90) {
			indexArray = append(indexArray, int(chara[i]))
		} else {
			indexArray = append(indexArray,int(-1))
		}
	}
	for j := 0; j < 26; j++ {
		for i := 0; i < len(indexArray) - 1; i++ {
			if index >= len(indexArray) {
				break
			}
			if indexArray[index] == -1 {
				output += string(chara[index])
				index += 1
				i--
			} else if (indexArray[i] - 97) == j || (indexArray[i] - 65) == j {
				output += string(chara[i])
				index += 1
			}
		}
	}

	return output
}

func sortCharaBucket(chara string) (string) {
	var output string
	var bucket [26]string
	var indexArray []int
	var indexPer int = 0
	var indexBucket int = 0

	for i := 0; i < len(chara) - 1; i++ {
		if chara[i] >= 65 && chara[i] <= 90 {
			bucket[chara[i] - 65] += string(chara[i])
			indexArray = append(indexArray, 1)
		} else if chara[i] >= 97 && chara[i] <= 122 {
			bucket[chara[i] - 97] += string(chara[i])
			indexArray = append(indexArray, 1)
		} else {
			indexArray = append(indexArray, -1)
		}
	}

	for i := 0; i < len(chara) - 1; i++ {
		if indexArray[i] == 1 {
			for ;indexPer >= len(bucket[indexBucket]) && indexBucket < 26; indexBucket++ {
				indexPer = 0
			}
			output += string(bucket[indexBucket][indexPer])
			indexPer++
		} else {
			output += string(chara[i])
		}
	}

	return output
}