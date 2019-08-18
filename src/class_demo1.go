package main

import "fmt"

func main () {
	var a Integer = 1
	a.Add(2)
	fmt.Println(a)	// 1

	a.AddInPointer(2)
	fmt.Println(a)	// 3
}

type Integer int

func (a Integer) Add (b Integer) {
	a += b
}

func (a *Integer) AddInPointer(b Integer) {
	*a += b
}