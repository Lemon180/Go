package main

import "fmt"

/*
type 关键字
1、定义结构体
2、定义接口
3、定义类型别名
4、类型定义
5、类型判断
*/
func main() {

	// 定义别名 定义别名的时候需要使用 = 号
	type Myint = int
	var i Myint = 1         // 在编译的时候会把 Myint 替换为 int
	fmt.Printf("%T\r\n", i) // 输出类型还是 int

	// 自定义类型，需要基于现有的类型，定义类型的时候不要使用 = 号。这两个类型 Myfloat32 和 float32 是不一样的
	type Myfloat32 float32
	var a Myfloat32 = 3.1415926
	var b float32 = 3.1415926 // 两个类型是不一样的不能相加,需要类型转换

	fmt.Printf("%d", float32(a)+b) // main.Myfloat32
}
