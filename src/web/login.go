package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



func login(w http.ResponseWriter, r *http.Request) {
	//var TemplatesFiles = []string{
	//	"web/login.gtpl",
	//}

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {

		t, _ := template.ParseFiles("src/web/login.gtpl")
		//t, _ := template.ParseFiles(TemplatesFiles...)
		t.Execute(w, nil)
		//fmt.Println(t)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func TestLogin() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
