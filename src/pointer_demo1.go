package main

import (
	"fmt"
)

type Person struct {
	age int
}

func (p Person) getAge() int { // 接收值变量
	return p.age
}

func (p *Person) growAge() { // 接收指针变量
	p.age += 1
}

func main() {
	p1 := Person{24}         // p1为值变量
	fmt.Println(p1.getAge()) // 正常的值传递，方法使用调用的副本
	p1.growAge()             // 使用值的引用来调用此方法，(&p1).growAge()
	fmt.Println(p1.getAge())

	p2 := &Person{10}        // p2为指针变量
	fmt.Println(p2.getAge()) // 指针被解引用为值，(*p2).getAge()
	p2.growAge()             // 指针传参，拷贝了一份指针
	fmt.Println(p2.age)

	fmt.Printf("p2: %T, %v", p2, p2)
}
