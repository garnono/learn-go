package web

import (
	"log"
	"net/http"
)

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
func TestWeb() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/html", html)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(":ListenAndServe:", err)
	}
}