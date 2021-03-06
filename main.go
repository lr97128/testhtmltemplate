package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("read html file failed, err:", err)
		return
	}
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./auth.html")
	if err != nil {
		fmt.Println("read html file failed, err:", err)
		return
	}
	// var str string
	r.ParseForm() //先要解析form表单
	username := r.Form.Get("username")
	passwd := r.Form.Get("passwd")
	fmt.Println("username:", username, " passwd:", passwd)
	var user = User{
		Name:     username,
		Password: passwd,
	}
	// if username == "liurui" && passwd == "1234" {
	// 	str = "认证通过"
	// 	t.Execute(w, str)
	// } else {
	// 	str = "认证失败"
	// }
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
