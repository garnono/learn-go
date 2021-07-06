package basic

import (
	"bytes"
	"fmt"
	"strings"
)

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

/**
字符串 拼接的方法
1、+ 号
2、sprintf 函数
3、Join 函数
4、buffer.WriteString 函数
5、buffer.Builder（据说官方推荐）
*/

func JoinString() {
	fmt.Println("test join string start --------------")

	str1 := "abc"
	str2 := "123"

	strJoin1 := str1 + str2
	fmt.Println("+ :", strJoin1)

	fmt.Sprintln("%s%s", str1, str2)

	var strSlice []string = []string{str1, str2}
	strJoin3 := strings.Join(strSlice, "")
	fmt.Println("Join:", strJoin3)

	var bt bytes.Buffer
	bt.WriteString(str1)
	bt.WriteString(str2)
	strJoin4 := bt.String()
	fmt.Println("buffer.writerString:", strJoin4)

	var bd strings.Builder
	bd.WriteString(str1)
	bd.WriteString(str2)
	strJoin5 := bd.String()
	fmt.Println("buffer.Build:", strJoin5)

	fmt.Println("test join string end --------------")
}
