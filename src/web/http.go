package web

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	fmt.Fprintf(w, "hello")
}

func TestHttp() {
	// 设置路由
	http.HandleFunc("/", sayHelloName)
	// 监听
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(":ListenAndServe:", err)
	}
}
