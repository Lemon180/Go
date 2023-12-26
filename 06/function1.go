package main

import "fmt"

/*
函数一等公民
*/

// 函数当做变量
func add(a int, b int) (sum int, err error) {
	sum = a + b
	return
}

// 函数返回函数
func cal(ops string) func() {
	switch ops {
	case "+":
		return func() {
			fmt.Println("加法")
		}
	case "-":
		return func() {
			fmt.Println("减法")
		}
	case "*":
		return func() {
			fmt.Println("乘法")
		}
	case "/":
		return func() {
			fmt.Println("除法")
		}
	default:
		return func() {
			fmt.Println("错误的操作符")
		}
	}
}

// 函数当做参数使用
func cal2(myfunc func(itme ...int) int) int {
	return myfunc()
}
func main() {

	// 函数当做变量名使用
	funcVar := add
	var sum, _ = funcVar(1, 2)
	fmt.Println(sum)

	// 函数返回函数
	cal("+")()

	// 匿名函数当做参数使用
	mysum := cal2(func(item ...int) int {
		sum := 0
		for _, value := range item {
			sum += value
		}
		return sum
	})
	fmt.Println(mysum)
}
