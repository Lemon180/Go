package main

import (
	"container/list"
	"fmt"
)

/*
*
list 在Go语言中是链表
1、在频繁插入和删除的时候可以使用 list
*/
func main() {

	// 创建list对象
	var myList list.List
	//  向 list 前添加元素
	myList.PushFront("第一个对象")
	// 向 list 后添加元素
	myList.PushBack("第二个对象")

	//fmt.Println(myList)
	//遍历 list 正向遍历
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// 反向遍历 list
	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	// 删除list中的元素
	myList.Remove(myList.Front())
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
