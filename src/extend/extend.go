package extend

import (
	"extend/cli"
	"strings"
)

// 可执行的方法映射关系
var command = map[string]func(){
	"args": cli.TestArgs,
}

// 通过映射，动态调用方法
func call(com string) {
	if function, ok := command[com]; ok {
		function()
	}
}

// extend 内容测试入口
func TestExtend(action string) {
	call(action)
}

func GetCommandKeyStr() string {
	var keys []string

	for ak, _ := range command {
		keys = append(keys, ak)
	}
	return strings.Join(keys, "|")
}