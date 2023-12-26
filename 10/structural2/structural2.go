package main

import "fmt"

/*
结构体嵌套 就相当与 java中的继承
*/

type Person struct {
	name string
	age  int
}
type Student struct {
	//p      Person // 第一种方式嵌套Person结构体
	Person // 第二种方式嵌套Person结构体（匿名嵌套）
	score  float32
}

func main() {
	s := Student{
		Person{name: "张三", age: 12},
		95.6,
	}
	fmt.Println(s.name, s.age, s.score)
}
