package main

import "fmt"

/*
函数闭包
闭包（Closure）是指一个函数包含了它外部作用域中的变量，即使在外部作用域结束后，这些变量依然可以被内部函数访问和修改。
闭包使得函数可以“记住”外部作用域的状态，这种状态在函数调用之间是保持的。
闭包的核心概念是函数内部可以引用外部作用域的变量，即使在函数内部外部作用域已经结束。


闭包的特点
函数可以在定义的作用域之外被调用，仍然可以访问外部作用域的变量。
外部作用域中的变量不会被销毁，直到闭包不再引用它们。
多个闭包可以共享同一个外部作用域的变量。
*/

func auoIncrement() func() int {
	local := 0
	return func() int {
		local += 1 // local 变量是在 auoIncrement 函数中的变量，这种方式是闭包 可以访问上一个函数中的变量
		return local
	}
}
func main() {
	nextFunc := auoIncrement() // 这里返回的是一个函数 nextFunc 是一个函数
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
}
