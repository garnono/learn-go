package learn

import "fmt"

func TestVar() {
	PrintStart("var")

	/*
		三种声明方式：
		1、常规：var identifier type
		2、省略类型：var v_name = value
		3、内部：v_name := value
	*/

	/*
		多变量声明

		var a,b int

		var (	// 全局
			a int
			b string
		)

		a,b := 1, "string"

	*/

	/*
		值类型：基本类型 int、float、string、bool；赋值是值的拷贝
		引用类型：
	*/

	i1 := 123
	i2 := i1
	fmt.Println("i 1", i1, "ptr", &i1)
	fmt.Println("i 2", i2, "ptr", &i2)

	a1 := []int{1, 2}
	a2 := a1
	a2[1] = 3
	fmt.Println("a 1", a1, "ptr", &a1)
	fmt.Println("a 2", a2, "ptr", &a2)

}
