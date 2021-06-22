package learn

import (
	"errors"
	"fmt"
)

func TestInterface() {
	PrintStart("interface")

	/**
	定义：
	type interface_name interface {
		method_name1 [return_type]
		method_name2 [return_type]
	}
	*/

	p1 := IPhone{}
	p1.name = "iPhone"
	_, err := p1.call()
	fmt.Println("error:", err)

	// 存储的值——任意类型的值
	var mobile Phone

	iphone := IPhone{name: "iphone"}
	android := Android{name: "android"}

	mobile = iphone
	mobile = android

	// 判断类型
	if _, ok := element.(Android); ok {
		fmt.Println("is android")
	}

	switch v := element.(type) {
	case IPhone:
		fmt.Println("is iphone", v)
	case Android:
		fmt.Println("is andorid", v)
	}

}

type Phone interface {
	call() (int, error)
	send()
}

type IPhone struct {
	name string
}

type Android struct {
	name string
}

func (p IPhone) call() (int, error) {
	fmt.Println(p.name, "to call")
	return 1, errors.New("nothing")
}

func (p IPhone) send() {
	fmt.Println("send message...")
}

func (p Android) call() (int, error) {
	fmt.Println(p.name, "to call")
	return 1, errors.New("nothing")
}

func (p Android) send() {
	fmt.Println("send message...")
}
