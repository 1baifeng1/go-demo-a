package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a Integer = 1
	a.Add(2)
	fmt.Println(a) // 1

	a.AddInPointer(2)
	fmt.Println(a) // 3
	fmt.Println("///////////////////////////")
	fmt.Println("sizeof(T1{})", unsafe.Sizeof(T1{}))

	age1 := T2{3}
	//var age2 = new(T2)
	//age2.age = 3
	age2 := T2{3}
	// 只有同类型的结构体能比较，比较的是值不是结构体地址
	fmt.Println("age1 == age2 ? ", age1 == age2)
	fmt.Printf("age1 : %p, age2 : %p \n", &age1, &age2)
	//age3 := T3{3}
	//fmt.Println("age2 == age3 ? ", age2 == age3)
	// invalid operation: age2 == age3 (mismatched types T2 and T3)

}

type Integer int

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) AddInPointer(b Integer) {
	*a += b
}

//////////////////////////
type T1 struct{}

type T2 struct {
	age int
}
type T3 struct {
	age int
}
