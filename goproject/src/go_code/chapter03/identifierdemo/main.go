package main

import "fmt"

//演示go中标识符的使用
func main() {
	//go中严格区分大小写
	//go中认为 num 和 Num 是不同的变量
	var num int = 20
	var Num int = 10
	fmt.Printf("num=%v Num=%v\n", num, Num)
}
