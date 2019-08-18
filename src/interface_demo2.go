package main

import "fmt"

type IGreeting interface{
	sayHello()
}

// 可以通过接口实现多态
func sayHello(greetings IGreeting) {
	greetings.sayHello()
}

type Go struct {
	name string
	compileType string
}

type Python struct {
	name string
	compileType string
	version float32
}

func (g Go) sayHello() {
	fmt.Println(g.name, g.compileType)
}

func (p Python) sayHello() {
	fmt.Println(p.name, p.compileType, p.version)
}

func main() {
	g := Go{"Go", "static"}
	g.sayHello()

	py := Python{"Python", "dynamic", 2.7}
	sayHello(py)
}