package main

import "fmt"

/*
指针
值传递
*/
type Person struct {
	name string
}

func ChangeName(p Person) {
	p.name = "李四"
}
func main() {

	p := Person{name: "张三"}
	ChangeName(p) // 这里修改结构体中的值 张三 改为 李四。发现值输出的依然是张三,函数并没有改变结构体中的值
	fmt.Println(p.name)
}
