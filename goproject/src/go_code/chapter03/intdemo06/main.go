package main

import (
	"fmt"
	"unsafe"
)

// 演示go中整数类型使用
func main() {
	var i int = 1
	fmt.Println("i=", i)

	//测试一下int8的范围 -128~127
	//其他的 int16，int32，int64，类推。。。
	var j int8 = 127
	fmt.Println("j=", j)

	//测试一下uint8的范围 0~255
	//其他的 uint16，int32，int64，类推。。。
	var k uint8 = 255
	fmt.Println("k=", k)

	//int,uint,rune,byte的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	//整型的使用细节
	var n1 = 100 //?n1是什么类型
	//fmt.Printf()可以用于做格式化输出
	fmt.Printf("n1的类型%T\n", n1)

	//如何在程序查看某个变量的占用字节大小和数据类型（使用较多）
	var n2 int64 = 10
	//unsafe.Sizeof(n1)是unsafe包的一个函数，可以返回n2变量占用的字节数
	fmt.Printf("n2的类型%T n2占用的字节数是%d\n", n2, unsafe.Sizeof(n2))

	//go程序中整型变量在使用时，遵守保小不保大的原则，
	//即：在保证程序正确运行下，尽量使用占用空间下的数据类型
	var age byte = 90
	fmt.Println("年龄：", age)
}
