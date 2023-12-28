package main

import "fmt"

/*
多结构体 实现 多接口
*/

// 定义接口1
type WriterFile interface {
	fileWriter(string) error
}

// 定义接口2
type WriterClose interface {
	fileClose() error
}

// 定义结构体1
type Files struct{}

// 结构体1实现接口1
func (f *Files) fileClose() error {
	fmt.Println("Files-WriterClose")
	return nil
}

// 结构体1实现接口2
func (f *Files) fileWriter(s string) error {
	fmt.Println("Files-WriterFile")
	return nil
}

// 定义结构体2
type Database struct{}

// 结构体2实现接口1
func (d *Database) fileClose() error {
	fmt.Println("Database-WriterClose")
	return nil
}

// 结构体2实现接口2
func (d *Database) fileWriter(s string) error {
	fmt.Println("Database-WriterFile")
	return nil
}

func main() {
	// 结构体1实现
	var mf WriterFile = &Files{}
	var mc WriterClose = &Files{}
	fmt.Println(mf.fileWriter("读文件"))
	fmt.Println(mc.fileClose())
	// 结构体2实现
	var dbmf WriterFile = &Database{}
	var dbmc WriterClose = &Database{}
	fmt.Println(dbmf.fileWriter("读数据库"))
	fmt.Println(dbmc.fileClose())

}
