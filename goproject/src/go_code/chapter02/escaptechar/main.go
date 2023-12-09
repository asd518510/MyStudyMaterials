package main

import "fmt" //fmt包中提供格式化，输出，输入的函数。

// 这是一个main函数，是程序入口
func main() {
	//演示转义字符的使用 \t
	fmt.Println("tom\tjack")

	//如果希望一次性注释 ctrl + / 第一次表示注释，第二次表示取消注释
	fmt.Println("hello\nworld")
	fmt.Println("E:\\学习库\\MyStudyMaterials\\GOPROJECT\\src\\go_code")
	fmt.Println("tom说\"i love you\"")

	// \r 回车,从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

}
