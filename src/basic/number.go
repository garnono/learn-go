package basic

import "fmt"

func TestNumber() {
	PrintStart("number")

	// 特点：原生支持复数，其中位的运算采用补码

	// 声明

	/*
		整型：int，包括有符号、无符号两大类；根据长度分为：8、16、32、64；
		另外：byte =》 uint8
		rune => int32
		uint 32 or 64位
		int => uint
		uintptr 无符号整型，用于存放一个指针
	*/
	var i1 int
	var i2 int8
	var i3 int16
	var i4 int32
	var i5 int64

	var ui1 uint
	var ui2 uint8
	var ui3 uint16
	var ui4 uint32
	var ui5 uint64

	fmt.Println("int ", i1)
	fmt.Println("int 8", i2)
	fmt.Println("int 16", i3)
	fmt.Println("int 32", i4)
	fmt.Println("int 64", i5)

	fmt.Println("uint ", ui1)
	fmt.Println("uint 8", ui2)
	fmt.Println("uint 16", ui3)
	fmt.Println("uint 32", ui4)
	fmt.Println("uint 64", ui5)

	var b byte
	var r rune
	var uip uintptr

	fmt.Println("byte ", b)
	fmt.Println("rune", r)
	fmt.Println("uintptr", uip)

	/*
		浮点型：float，包括float32，float64，complex64，complex128
	*/

	var f1 float32
	var f2 float64
	var c1 complex64
	var c2 complex128

	fmt.Println("float 32", f1)
	fmt.Println("float 64", f2)
	fmt.Println("complex 64", c1)
	fmt.Println("complex 128", c2)

	// 默认值
	// 声明 & 赋值
	// 注意：
}
