package main

import "fmt"

/*
type 定义结构体
*/
func main() {

	// 定义结构体 可以理解为java中的实体类
	type Person struct {
		name    string
		age     int
		height  int
		address string
	}

	// 初始化结构体
	//1、 按顺序初始化
	p1 := Person{"zhangsan", 18, 170, "beijing"}
	//2、按字段初始化
	p2 := Person{
		name:    "lisi",
		age:     20,
		height:  160,
		address: "shanghai",
	}
	fmt.Println("P2=====", p2)

	// 使用 Person 类型的切片
	var Persons []Person
	// 可以将 Person 结构体初始化后放入切片
	Persons = append(Persons, p1)
	// 也可以在放入切片时将 Person 结构体初始化
	Persons = append(Persons, Person{name: "张三"})
	fmt.Println("Person=====", Persons)

	// 方法1、使用 Person 类型的切片 并初始化
	Person2 := []Person{
		{name: "王五"},
	}
	fmt.Println("Person2=====", Person2)
	//方法2、使用 Person 类型的切片 并初始化
	Person3 := []Person{
		{"zhangsan", 18, 170, "beijing"},
	}
	fmt.Println("Person3=====", Person3)

	// 结构体取值 和 赋值
	fmt.Println(Person3[0].name)
	Person3[0].name = "新值"
	fmt.Println(Person3[0].name)

	// 匿名结构体,在函数内部或者只用一次的时候,创建好并且初始化
	adders := struct {
		province string
		city     string
	}{
		province: "北京",
		city:     "北京",
	}
	fmt.Println(adders.city)
}
