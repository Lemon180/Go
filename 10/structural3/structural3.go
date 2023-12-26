package main

import "fmt"

/*
结构体定义方法

定义方法的语法格式：
		值传递 传递的是复制后的值
	func (接收者变量 接收者类型) 方法名(参数列表) (返回值列表){
	    方法体
	}
		引用传递 传递的是指针
	func (接收者变量 *接收者类型) 方法名(参数列表) (返回值列表){
	    方法体
	}
*/

type Person struct {
	name    string
	age     int
	address string
}

// 函数可以直接调用
func NewPerson(name string, age int, address string) {
	p := Person{name, age, address}
	fmt.Println(p)
}

// 结构体方法 是绑定在结构体上的，必须初始化结构体之后才能使用
// 这种是值传递
func (p Person) getName() string {
	return p.name
}

// 结构体方法 是绑定在结构体上的，必须初始化结构体之后才能使用
// 这种是引用传递
func (p *Person) getAge() int {
	return p.age
}

func main() {

	p := Person{"zhangsan", 18, "北京"}
	name := p.getName()
	age := p.getAge()
	fmt.Println(name, age)
	NewPerson("hahaha", 18, "上海")
}
