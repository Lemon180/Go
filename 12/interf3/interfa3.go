package main

/*
结构体嵌套接口
*/

type person interface {
	speak()
}

type student struct {
	person // 嵌套接口
}

func main() {

}
