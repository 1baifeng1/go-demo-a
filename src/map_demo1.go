package main

import "fmt"

func say(text string) {
	fmt.Println(text)
}

func map_func() {
	// 创建map时，value可以是函数
	var funcMap = make(map[string]func(string))
	funcMap["say"] = say
	funcMap["say"]("hello world.")
}

func main() {
	map_func()
}

