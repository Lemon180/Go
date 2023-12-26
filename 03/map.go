package main

import "fmt"

/*
*
map数据结构
1、map 是一种无序的元素集合 由key 和value 组成
2、想要往Map中放值，必须要初始化Map
3、Map不是线程安全的，所以不要在多协程中使用Map
*/
func main() {

	// 1、声明并初始化Map
	var courseMap = map[string]string{
		"go":   "go语言工程师",
		"java": "java语言工程师",
		"php":  "php语言工程师",
	}
	// 从Map中取值
	fmt.Println(courseMap["go"])
	// 放值
	courseMap["mysql"] = "mysql工程师"
	fmt.Println(courseMap)

	// 2、初始化方法2
	// 在Go语言中 make 函数用来初始化 slice Map channel
	var courseMap2 = make(map[string]string)
	courseMap2["google"] = "google1工程师"
	courseMap2["google1"] = "google2工程师"
	courseMap2["google2"] = "google3工程师"
	courseMap2["google3"] = "google4工程师"
	courseMap2["google4"] = "google5工程师"
	courseMap2["google5"] = "google6工程师"

	fmt.Println(courseMap2)

	// 遍历Map
	for key, value := range courseMap2 {
		fmt.Println(key, value)
	}

	// 判断Map中是否存在
	if value, ok := courseMap2["google2"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("not")
	}
	//删除Map中的数据
	delete(courseMap2, "google4")
	fmt.Println(courseMap2)
}
