package main

import (
	"basic"
	"extend"
	"extend/cli"
	"flag"
	"fmt"
	"web"
)

func main() {
	fmt.Println("hello goland !")

	cli.TestOs()
	//cli.TestFlag()

	basicActionStr := basic.GetCommandKeyStr()
	extendActionStr := extend.GetCommandKeyStr()

	modeDesc := "mode，可选值：basic｜web｜extend"
	actionDesc := "action，可选值：当mode=basic时，" +
		"\n当mode=basic时，" + basicActionStr +
		"\n当mode=web时，action 无效" +
		"\n当mode=extend时，" + extendActionStr

	mode := flag.String("m", "", modeDesc)
	action := flag.String("a", "", actionDesc)
	flag.Parse()

	switch *mode {
	case "basic":
		basic.TestBasic(*action)
	case "web":
		web.TestWeb()
	}
}
