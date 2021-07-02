package basic

import (
	"fmt"
	"time"
)

func TestOther() {
	PrintStart("other")

	// 并发
	go say("hello")
	say("go")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
