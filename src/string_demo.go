package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 测试spilt的使用
func spilt_study() {
	str := "1 2 10 111 2"
	str_arr := strings.Split(str, " ")
	//fmt.Println(str_arr)
	for _, s := range str_arr {
		i, _ := strconv.Atoi(s)
		fmt.Println(i)
	}
}

// 引号
func quotes_study() {
	var a string = "你\n好呀"		// 双引号中支持转义，中文在utf-8中占3个字符
	fmt.Println(a, len(a))

	// 反单引号不支持转义，但是可以换行，会保留字符串原始的格式
	var b string = `start
    现在
我在\n哪里
？
stop`
	fmt.Println(b)

	var c rune = 'c'	// 输出c的ASCII值，英文字符在utf-8中占1个字节，中文在utf-8中占3字节
	var c2 rune = '中'	// 输出中的utf-8值
	fmt.Println(c, c2)
}

func main() {
	spilt_study()
	quotes_study()

}
