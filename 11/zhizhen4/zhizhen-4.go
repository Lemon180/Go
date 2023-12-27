package main

import "fmt"

/*
方法使用指针

	1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	2.指针变量的值是指针地址。
	3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/
type Person struct {
	name string
	age  int
}

// 定义结构体方法 使用指针类型
func (p *Person) SayHello() string {
	p.name = "Hello"
	return p.name
}

func main() {
	p := &Person{name: "Alice", age: 20}
	fmt.Println(p.SayHello())
}
