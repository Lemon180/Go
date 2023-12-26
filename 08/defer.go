package main

import "fmt"

/*
defer 在函数退出时释放资源 例如数据库连接 打开关闭文件 打开锁 释放锁
如果有多个defer 则按照后进先出的顺序执行
defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出
*/

func deferRes() (sum int) {
	defer func() {
		sum++
	}()
	return 10
}

func main() {
	fmt.Println(deferRes()) // 返回 11
}
