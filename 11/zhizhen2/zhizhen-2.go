package main

import "fmt"

/*
指针
引用传递

定义指针类型
var i = *int
p:= &Person{name: "张三"}
*/
type Person struct {
	name string
}

// *Person 传入的是一各指针类型的值
func ChangeName(p *Person) {
	p.name = "李四"
}
func main() {

	p := &Person{name: "张三"}
	ChangeName(p)       // &p 是取指针地址
	fmt.Println(p.name) // 输出发现结构体中的值被改变了
}
