package main

import "fmt"

/*
*
切片
切片底层也是数组
Go语言中的slice是值传递还是引用传递？是值传递，但是又呈现出引用传递的效果
1、在定义slice的时候可以指定指针位置、长度、和容量。
2、如果是参数或者函数 修改slice的时候，Go语言会复制出来一个新的结构体，和之前的结构体的指针位置指向同一个内存、长度和容量也指向同一个。
3、如果复制出来的slice发生了扩容，那么会重新分配内存，复制数据，指针位置和之前的结构体指针位置不变。
*/
func main() {
	// 定义切片 []
	var courses []string
	// 在 go 语言中使用 append() 函数向切片中追加元素
	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses)
	// 遍历切片
	for _, cours := range courses {
		fmt.Println(cours)
	}

	// 1、从数组初始化切片
	var courses1 = [5]string{"1", "2", "3", "4", "5"}
	fmt.Printf("%T\n", courses1)
	courses2 := courses1[0:2]
	fmt.Printf("%T\n", courses2)

	// 2、初始化切片
	courses3 := []string{"go", "grpc", "javascript", "gin"}
	fmt.Println(courses3)

	// 3、初始化切片 类型string 长度3 空间10
	courses4 := make([]string, 3, 10)
	fmt.Println(courses4)

	// 切片取值
	fmt.Println(courses3[2:3])
	fmt.Println(courses3[0:2])

	//将两个切片合并为一个切片 需要使用 append方法 参数使用...三个省略号将 courses6 打散添加到 courses5 中
	courses5 := []string{"java", "mysql", "redis"}
	courses6 := []string{"go", "grpc", "gin"}
	courses5 = append(courses5, courses6[0:]...)
	fmt.Println(courses5)

	// 删除切片中的元素
	courses7 := []string{"java", "mysql", "redis", "go", "grpc", "gin"}
	mySlice := append(courses7[:2], courses7[3:]...)
	fmt.Println(mySlice)

	// 复制切片 浅复制
	courses8 := []string{"1", "2", "3", "4", "5"}
	courses9 := courses8[:]
	courses9[0] = "java"
	fmt.Println(courses8, courses9)

	//复制切片 深复制
	courses10 := []string{"1", "2", "3", "4", "5"}
	var courses11 = make([]string, len(courses10))
	copy(courses11, courses10)
	courses11[2] = "999"
	fmt.Println(courses10, courses11)
}
