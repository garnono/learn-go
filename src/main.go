package main

import (
	"fmt"
	"learn"
)

func main() {
	fmt.Println("hello goland !")

	/*
		学习思路：
		1、结构
		    包声明				package main
		    引入包				import "fmt"
		    函数					func main() {
		    变量						var a int = 10
		    语句 & 表达式				fmt.Println("a:", a)
		    注释					}	// 函数结束——单行注释
		2、常量：运行时值不能改变
		3、变量：运行时值可以改变
		4、类型：对值的分类
			基本类型：bool/数字/字符串
			其他类型：数组/指针/结构体/切片/集合/channel
		5、运算符
		    算术运算符
				+	相加		A + B 输出结果 30
				-	相减		A - B 输出结果 -10
				*	相乘		A * B 输出结果 200
				/	相除		B / A 输出结果 2
				%	求余		B % A 输出结果 0
				++	自增		A++ 输出结果 11
				--	自减		A-- 输出结果 9
		    关系运算符
				==	检查两个值是否相等，如果相等返回 True 否则返回 False。			(A == B) 为 False
				!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。		(A != B) 为 True
				>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。		(A > B) 为 False
				<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。		(A < B) 为 True
				>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
				<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True
			逻辑运算符
				&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 	(A && B) 为 False
				||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。		(A || B) 为 True
				!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。		!(A && B) 为 True
		    位运算符
				&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 		(A & B) 结果为 12, 二进制为 0000 1100
				|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或			(A | B) 结果为 61, 二进制为 0011 1101
				^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。				(A ^ B) 结果为 49, 二进制为 0011 0001
				<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 	A << 2 结果为 240 ，二进制为 1111 0000
				>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 					A >> 2 结果为 15 ，二进制为 0000 1111
		    赋值运算符
				=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
				+=	相加后再赋值	C += A 等于 C = C + A
				-=	相减后再赋值	C -= A 等于 C = C - A
				*=	相乘后再赋值	C *= A 等于 C = C * A
				/=	相除后再赋值	C /= A 等于 C = C / A
				%=	求余后再赋值	C %= A 等于 C = C % A
				<<=	左移后赋值 	C <<= 2 等于 C = C << 2
				>>=	右移后赋值 	C >>= 2 等于 C = C >> 2
				&=	按位与后赋值	C &= 2 等于 C = C & 2
				^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
				|=	按位或后赋值	C |= 2 等于 C = C | 2
		    其他运算符
				&	返回变量存储地址	&a; 将给出变量的实际地址。
				*	指针变量。	*a; 是一个指针变量
			优先级（由高到低）
				5 	* / % << >> & &^
				4 	+ - | ^
				3 	== != < <= > >=
				2 	&&
				1 	||

		6、语句
			条件：if...else.../switch/select
			循环:for/break/continue/goto

		7、其他
			函数：作为参数/闭包/方法
			接口
			作用域：局部/全局/形参
			Range
			类型转换
			递归
			错误处理
			并发

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
	//类型：通道
	learn.TestChannel()

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

	// 其他
	learn.TestOther()

	c := Color{name: "test"}
	fmt.Println("c :", c)

	c.changeName("pointer")
	fmt.Println("c :", c)

	c.changeName2("other")
	fmt.Println("c :", c)
}

type Color struct {
	name string
}

func (c *Color) changeName(name string) {
	c.name = name
}

func (c Color) changeName2(name string) {
	c.name = name
}
