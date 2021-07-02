package web

// 可执行的方法映射关系
var command = map[string]func(){
	"http":  TestHttp,
	"login": TestLogin,
}

// 通过映射，动态调用方法
func call(com string) {
	if function, ok := command[com]; ok {
		function()
	}
}

// web 内容测试入口
func TestWeb(action string) {
	call(action)
}
