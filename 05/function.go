package main

import "fmt"

/*
函数
1、Go语言中函数可以当做变量
2、匿名函数 和 闭包
3、函数可以满足接口
4、Go语言中函数可以返回多个值
5、在Go语言中参数传递的时候都是值传递
*/

// 只返回一个值
func add(a int, b int) int {
	return a + b
}

// 返回两个值
func add2(a int, b int) (int, error) {
	return a + b, nil
}

// 定义返回值变量名
func add3(a int, b int) (sum int, err error) {
	sum = a + b
	return sum, err
}

// 可变参数
func add4(items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return
}
func main() {

	// 调用函数
	fmt.Println(add(1, 2))

	// 接收多个返回值
	sum, _ := add2(2, 2)
	fmt.Println(sum)

	// 调用可变参数函数
	sum2, _ := add4(1, 2, 3, 4, 5)
	fmt.Println(sum2)
}
