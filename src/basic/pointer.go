package basic

import "fmt"

func TestPointer() {
	PrintStart("pointer")

	// 特点：
	// 声明
	// 默认值
	// 声明 & 赋值

	structPointer()

	// 注意：
}

func structPointer() {

	var book1 = Book{1, "go", "bird", 123.12321}

	// 声明：var struct_pointer *Books
	var bookP1 *Book
	bookP1 = &book1

	fmt.Println("struct pointer book 1 pointer:", bookP1)

	fmt.Println("struct pointer as params")
	printBook(*bookP1, 1)

	fmt.Println("struct pointer book 1 id:", bookP1.id)
	fmt.Println("struct pointer book 1 id:", (*bookP1).id)
}
