package web

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func html(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("web/html.gtpl")
		t.Execute(w, nil)

	} else {
		r.ParseForm()

		fmt.Println("params:", r.Form)

		for k, v := range r.Form {
			fmt.Println(k, "=", v)
		}

		// 字符串不为空
		if len(r.Form.Get("username")) == 0 {
			fmt.Println("username is empty")
		}
		// 正则：是否为中文
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("username")); !m {
			fmt.Println("username is not chinese")
		}
		// 正则：是否为英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			fmt.Println("username is not english")
		}

		// 转化为int
		age, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("age is not number by atoi")
		}
		fmt.Println("age:", age)

		// 正则：是否为int
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			fmt.Println("age is not number by regxp")
		}

		// 正则：是否为email
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("email is err")
		}

		// 正则：是否为手机号
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			fmt.Println("mobile is err")
		}

		// 下拉框
		fruit := []string{"apple", "pear", "banana"}

		for _, v := range fruit {
			if v == r.Form.Get("fruit") {
				fmt.Println("fruit:", v)
			}
		}

		// 单选
		gender := []int{1, 2}
		for _, v := range gender {
			//if v == r.Form.Get("gender") {
			//	fmt.Println("gender:", v)
			//}
			fmt.Println("gender:", v)
		}

		// 复选框
		interest := []string{"football", "basketball", "tennis"}
		//a:=Slice_diff(r.Form["interest"], interest)
		//if a !=nil{
		//	fmt.Println()
		//}
		fmt.Println("interest:", interest)

	}
}
