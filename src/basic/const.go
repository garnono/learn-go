package basic

import "fmt"

func TestConst() {
	PrintStart("const")

	// 定义：const identifier [type] = value
	const PI0 float32 = 3.14
	// 省略类型
	const PI1 = 3.14

	// 多常量、可以用len(), cap(), unsafe.Sizeof()常量计算表达式
	const A, B = "abc", 12

	const (
		MALE   = 1
		FEMALE = 2
	)

	// iota 自增
	const (
		a1 = iota	// 默认从0开始
		a2			// 自增1
		a3 = 100	// 打断
		a4			// 数值同上
		a5 = iota	// 继续iota的自增，中段的也会包括在自增内
		a6			// 自增1
		a7 = 100 + iota	// 重置自增表达式
		a8			// 按上面的表达式自增
	)

	fmt.Println("iota", a1, a2, a3, a4, a5, a6, a7, a8)

}
