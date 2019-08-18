package main

import (
	"fmt"
)

type People struct {
	name string
}

type Student struct{
	name string
	number int
}

type IPeople interface{
	getName() string
}

type IStudent interface{
	getName() string
	getNumber() int
}

func (p People) getName() string {
	return p.name
}

func (s Student) getName() string{
	return s.name
}

func (s Student) getNumber() int{
	return s.number
}

// 接口查询
func main() {
	// 不能用变量少的对象赋值给类型大的接口变量
	//var p1 = new(People)
	//var s1 IStudent = p1 //Cannot use 'p1' (type *People) as type IStudent in assignment Type does not implement 'IStudent' as some methods are missing: getNumber() int less...

	s2 := new(Student)
	var p2 IPeople = s2

	if p3, ok := p2.(IPeople); ok {
		fmt.Println("p3: IPeople", p3)
	}
	if p4, ok := p2.(IStudent); ok {
		fmt.Println("p4: IStudent", p4)
	}
	if p5, ok := p2.(*People); ok {
		fmt.Println("p5: *People", p5)
	}
	if p6, ok := p2.(*Student); ok {
		fmt.Println("p6: *Student", p6)
	}

	switch p2.(type) {
	//switch t := p2.(type) {
	case IPeople:
		fmt.Println("p2 can be IPeople.")
		//fallthrough	// cann't fallthrougn in type switch
	case IStudent:
		fmt.Println("p2 can be IStudent.")
	case *People:
		fmt.Println("p2 can be People.")
	case *Student:
		fmt.Println("p2 can be Student.")
	default:
		fmt.Println("p2 can't be anyone!")
	}

}