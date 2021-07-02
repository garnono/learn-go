package basic

import "fmt"

func TestString() {
	PrintStart("string")

	// 特点：字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本
	// 长度固定、单字节组成、utf8的unicode文本
	// 声明
	var s1 string

	// 默认值
	printString(s1, 1)

	// 声明 & 赋值
	var s2 string = "abc"
	printString(s2, 2)

	var s3 = "123"
	printString(s3, 3)

	var s4 = "abc" +
		"def"
	printString(s4, 4)

	// 注意：
}

func printString(str string, num int) {
	fmt.Println("string", num, str, "len", len(str))
}
