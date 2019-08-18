package main

import "fmt"

type Code interface {
	code()
}

type Programmer struct {
	name string
}

func (p Programmer) code () {
	fmt.Println(p.name, "is coding.")
}

func main() {
	var c Code
	fmt.Println("c == nil : ", c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	var p *Programmer
	fmt.Println("p == nil : ", p == nil)		// 只要对象的动态值为nil，这个对象就是nil
	fmt.Printf("p: %T, %v\r\n", p, p)

	c = p
	fmt.Println("c == nil : ", c == nil)		// 只有接口的动态类型和动态值都为nil时，该接口为nil
	fmt.Printf("c: %T, %v", c, c)
}
