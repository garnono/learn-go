package main

import (
	"fmt"
	"learn"
)


func main()  {
	fmt.Println("hello goland !")

	// 注释——单行
	/*
	 注释——多行
	 */

	// 常量
	learn.TestConst()
	// 变量
	learn.TestVar()

	// 类型：bool
	learn.TestBool()
	// 类型：数字
	learn.TestNumber()
	//类型：字符串
	learn.TestString()
	//类型：数组
	learn.TestArray()
	//类型：结构体
	learn.TestStruct()
	//类型：分片
	learn.TestSlice()
	//类型：指针
	learn.TestPointer()

	// 运算符

	// 语句
	learn.TestSentence()
	// 函数
	learn.TestFunc()
	// 作用域

	// 递归

	// 类型转换

	// 接口
	learn.TestInterface()
	// 错误处理
}
