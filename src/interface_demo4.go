package main

import (
	"fmt"
	"unsafe"
)

type iface struct {
	itab, data uintptr		// 自定义的两个指针
}

func main() {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	// 将abc在内存中的内容强制解释成自定义的iface
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	// 就可以打印出动态类型和动态值的地址了
	fmt.Println("ia=",ia, " ib=", ib, " ic=", ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))

	/*
	ia= {0 0}  ib= {4825088 0}  ic= {4825088 824634195632}	// b和c的动态类型一致，都是*int
	5
	 */
}