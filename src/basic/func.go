package basic

import "fmt"

func TestFunc() {
	PrintStart("func")

	// 值传递
	a, b := 1, 2
	fmt.Println("before swap a=", a, "b=", b)
	a, b = swap(a, b)
	fmt.Println("after swap a=", a, "b=", b)

	// 引用传递
	fmt.Println("after swap a=", a, "b=", b, "a ptr", &a, "b ptr", &b)
	swap2(&a, &b)
	fmt.Println("after swap a=", a, "b=", b, "a ptr", &a, "b ptr", &b)

	// 作为参数
	f := func() int {
		return 12
	}

	f1 := tFunc(f)
	fmt.Println("func", f1)

	// 闭包
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	// 方法
	b1 := Book{id: 1}
	b1.read()

	b2 := Book{id: 2, author: "null"}
	fmt.Println("book2", b2)
	b2.setAuthor("hahahahahah")
	fmt.Println("book2", b2)
	b2.setAuthorByPoint("test")
	fmt.Println("book2", b2)

	// 继承
	student1 := Student{}
	student1.sayHi()

	// 重写
	teacher1 := Teacher{}
	teacher1.sayHi()
}

/*
声明：
func function_name( [parameter list] ) [return_types]{
   函数体
}
*/

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func tFunc(f func() int) string {
	if f() == 12 {
		return "hehe"
	} else {
		return "er..."
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func (book Book) read() {
	fmt.Println("book for read")
}

func (book *Book) setAuthorByPoint(author string) {
	book.author = author
}

func (book Book) setAuthor(author string) {
	book.author = author
}

type Human struct {
	name string
	sex  int
	age  int
}

type Student struct {
	Human  // 匿名字段
	school string
}

type Teacher struct {
	Human
	school string
}

func (h Human) sayHi() {
	fmt.Println("human hi...")
}

func (t Teacher) sayHi() {
	fmt.Println("teacher hi...")
}
