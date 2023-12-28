package main

import "fmt"

/*
定义接口
在 Go语言中 结构体和接口是分开的没有任何关系 结构体可以实现任意接口
*/

// 定义接口
type Duck interface {
	Gaga() string
	walk() string
}

// 定义结构体
type psDuck struct {
	name string
}

// 使用结构体实现接口方法
func (p *psDuck) Gaga() string {
	p.name = "大鸭子"
	return p.name + "Gaga"
}

func (p *psDuck) walk() string {
	return p.name + "walk"
}

func main() {

	var d Duck = &psDuck{name: "小黄鸭"}
	fmt.Println(d.Gaga())
	fmt.Println(d.walk())
}
