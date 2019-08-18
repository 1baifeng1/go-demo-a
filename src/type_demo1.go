package main

import "fmt"

func main() {
	var i int = 9

	var f float32
	f = float32(i)			// go中的类型转换都必须是强制类型转换
	fmt.Printf("f: %T, %v\r\n", f, f)

	f = 10.8
	a := int(f)
	fmt.Printf("a: %T, %v\r\n", a, a)

}
