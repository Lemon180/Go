package main

import "fmt"

/*
指针初始化
1. 指针变量声明后没有初始化，默认赋值为nil
2. 指针变量初始化后，如果没有赋值，则默认和声明时赋值一样
3. 指针变量初始化后，可以通过“.”来访问指针指向的内容
4. 指针变量初始化后，可以通过“->”来访问指针指向的内容
5. 指针变量初始化后，可以通过“*”来取消指针的解引用
6. 指针变量初始化后，可以通过“&”来获取指针的地址
7. 指针变量初始化后，可以通过“new”来创建指针
8. 指针变量初始化后，可以通过“make”来创建指针
9. 指针变量初始化后，可以通过“delete”来删除指针
10. 指针变量初始化后，可以通过“copy”来复制指针
11. 指针变量初始化后，可以通过“swap”来交换指针的值
12. 指针变量初始化后，可以通过“len”来获取指针的长度
*/
type Person struct {
	name string
	age  int
}

func main() {
	//第一种初始化方法
	P1 := &Person{name: "Alice", age: 30}
	fmt.Println(P1.name)
	// 第二种初始化方法
	var P2 Person
	P := &P2
	fmt.Println(P.name)
	// 第三种初始化方法
	var P3 = new(Person)
	fmt.Printf("%p", P3)
}
