package main

import (
	"basic"
	"flag"
	"fmt"
	"web"
)

func main() {
	fmt.Println("hello goland !")

	str := basic.GetCommandKeyStr()
	fmt.Println(str)


	modeDesc := "mode，可选值：basic｜web"
	actionDesc := "action，可选值：当mode=basic时，\n当mode=web时，"

	mode := flag.String("m", "", modeDesc)
	action := flag.String("a", "", actionDesc)
	flag.Parse()

	switch *mode {
	case "basic":
		basic.TestBasic(*action)
	case "web":
		web.TestWeb(*action)
	}
}
