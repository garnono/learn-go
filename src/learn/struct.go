package learn

import "fmt"

/*
	声明：
	type struct_variable_type struct {
	   member definition;
	   member definition;
	   ...
	   member definition;
	}
*/
type Book struct {
	id     int
	title  string
	author string
	price  float32
}

func TestStruct() {
	PrintStart("struct")

	// 特点：类型可不同，结构固定(不能len)，内容可修改

	// 声明 struct
	var book1 Book

	// 默认值：不是nil
	printBook(book1, 1)
	//if book == nil {
	//	fmt.Println("book is nil")
	//}

	book1.id = 1
	book1.title = "go"
	book1.author = "bird"
	book1.price = 3.123
	printBook(book1, 1)

	// 声明 & 赋值
	var book2 = Book{2, "php", "bee", 4.21312}
	printBook(book2, 2)

	book3 := Book{3, "python", "cook", 2.1231}
	printBook(book3, 3)

	// 注意：

}

func printBook(book Book, num int) {
	fmt.Println("book", num, "value :", book)
	fmt.Println("book", num, "id:", book.id)
	fmt.Println("book", num, "title:", book.title)
	fmt.Println("book", num, "author:", book.author)
	fmt.Println("book", num, "price:", book.price)
}
