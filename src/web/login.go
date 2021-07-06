package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// 解析模版
		//cli目录需要注意，如果在项目根目录下允许，则为src/web/login.gtpl；如果是src下，则此处需要去掉src目录
		t, _ := template.ParseFiles("web/login.gtpl")
		// 渲染 & 发送
		t.Execute(w, nil)
	} else {
		// 解析参数
		r.ParseForm()

		//username := r.Form["username"]

		// 获取的是 slice 格式？
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// 单个的值
		fmt.Println("username2:", r.Form.Get("username"))
	}
}

func TestLogin() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
