package main

import "fmt"

/*
*
数组
*/
func main() {

	// 数组定义
	// 在 go 语言中 两个数组的长度不一致 两个数组的类型也不一致
	var course1 [5]string
	var course2 [10]string

	course1[0] = "你"
	course1[1] = "好"
	course1[2] = "吗"
	course1[3] = "世"
	course1[4] = "界"
	// 一个是 [5]string类型 一个是[10]string类型 所以两个是不同的类型
	fmt.Printf("%T\t,%T\n", course1, course2)

	// 数组遍历
	for index, value := range course1 {
		fmt.Println(index, value)
	}
	//数组初始化1
	course3 := [4]string{"1", "2", "3", "4"}
	for index, value := range course3 {
		fmt.Println(index, value)
	}
	//数组初始化2  ... 代表任意长度的数组 根据元素个数
	course4 := [...]string{"A", "B", "C", "D"}
	for index, value := range course4 {
		fmt.Println(index, value)
	}

	//数组比较 在go语言中 数组可以直接比较，但是两个数组的长度和类型必须一致，go语言会按顺序比较两个数组中的值
	if course3 == course4 {
		fmt.Println("两个数组一致")
	} else {
		fmt.Println("两个数组不致")
	}
	//多维数组
	var course5 = [2][4]string{}
	course5[0] = [4]string{"语文", "数学", "英语", "政治"}
	course5[1] = [4]string{"历史", "生物", "地理", "科学"}
	//遍历多维数组
	for _, strings := range course5 {
		//遍历一维数组
		for _, value := range strings {
			//打印元素
			fmt.Print(value)
		}
		//打印换行
		fmt.Println()
	}
}
