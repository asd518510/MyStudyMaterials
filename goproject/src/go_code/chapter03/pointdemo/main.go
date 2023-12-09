package main

import (
	"fmt"
)

// 演示go中指针类型
func main() {
	//基本数据类型在内存的布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址是", &i)

	//下面的var ptr *int = &i
	//1.ptr 是一个指针变量
	//2.ptr的类型 *int
	//3.ptr 本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)

	var num int = 9
	fmt.Printf("num adderss=%v\n", &num)

	var ptr1 *int
	ptr1 = &num
	*ptr1 = 10
	fmt.Printf("num =%v\n", num)
}
