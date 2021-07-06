package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func TestArgs() {

}

func TestOs() {
	fmt.Println("os 方式接受打印 cli 参数 ----- start ------")

	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	fmt.Println("os 方式接受打印 cli 参数 ----- end ------")
}

func TestFlag() {
	fmt.Println("flag 包方式接受打印 cli 参数 ----- start ------")

	t1 := flag.Bool("t1", false, "t1 bool 类型")
	t2 := flag.Int("t2", 0, "t2 int 类型")
	// 必须先解析，才能使用
	flag.Parse()
	fmt.Println("t1:", *t1)
	fmt.Println("t2:", *t2)

	fmt.Println("flag 包方式接受打印 cli 参数 ----- end ------")
}

func TestUrfave(){
	// https://blog.csdn.net/TurkeyCock/article/details/80359654

	/*
	NAME:
	   GoTest - hello world

	USAGE:
	   GoTest [global options] command [command options] [arguments...]

	VERSION:
	   1.2.3

	COMMANDS:
	     help, h  Shows a list of commands or help for one command
	   arithmetic:
	     add, a  calc 1+1
	     sub, s  calc 5-3
	   database:
	     db  database operations

	GLOBAL OPTIONS:
	   --lang FILE, -l FILE    read from FILE (default: "english")
	   --port value, -p value  listening port (default: 8000)
	   --help, -h              Help!Help!
	   --print-version, -v     print version
	 */
}
