package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main()  {
	var x float64 = 3.14
	fmt.Println("type of x:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	// Both Value.Type() and Value.Kind() can get the type
	fmt.Println("type of v:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	var t1 T
	vt1 := reflect.ValueOf(t1)
	// 对于结构体，Type()可以获取结构体信息，Kind()只能拿到"struct"
	fmt.Println("vt1 type is", vt1.Type(), "\tvt1 kind is", vt1.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println("v.CanSet() is", v.CanSet())
	// 1通过变量的指针
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("p.CanSet() is", p.CanSet())
	// 2拿到value的值
	v1 := p.Elem()
	fmt.Println("type if v1:", v1.Type())
	fmt.Println("v1.CanSet() is", v1.CanSet())
	// 3就可以修改原变量
	if v1.CanSet(){
		v1.SetFloat(4.13)
	}
	fmt.Println("v1 =", v1, "\tx =", x)

	fmt.Println()
	reflectT()
}

func reflectT() {
	t := T{123, "int123"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		// Interface()可以获取value的值，但是interface{}对象
		fmt.Printf("%d: %s %s = %v\n",
			i, typeOfT.Field(i).Name,f.Type(), f.Interface())
	}

	fmt.Println("")
}