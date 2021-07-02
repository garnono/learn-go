package basic

import "fmt"

func TestBool() {
	PrintStart("bool")

	// 特点：true false
	// 声明
	var b bool

	// 默认值
	fmt.Println("bool is ", b)

	// 声明 & 赋值
	var b1 bool = true
	fmt.Println("b1 is", b1)

	var b2 = false
	fmt.Println("b2 is", b2)

	// 注意：
}
