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

}

type Phone interface {
	call() (int, error)
	send()
}

type IPhone struct {
	name string
}

func (p IPhone) call() (int, error) {
	fmt.Println(p.name, "to call")
	return 1, errors.New("nothing")
}
