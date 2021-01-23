package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)


func loginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/login.html")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		//不好用:fmt.Println("username:", r.Form["username"])
		fmt.Println("username:", r.FormValue("username"))
		//不好用:fmt.Println("password:", r.Form["password"])
		fmt.Println("password:", r.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/login", loginFunc)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
http://localhost:9090/login
http://192.168.219.134:9090/login
*/

/*

Tips:
[2]Request本身也提供了FormValue()函数来获取用户提交的参数。
如r.Form["username"]也可写成r.FormValue("username")。
[2]调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。
[3]r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
*/

/*
[1]通过上面的代码我们可以看出获取请求方法是通过r.Method来完成的，这是个字符串类型的变量，返回GET, POST, PUT等method信息。
[2]login函数中我们根据r.Method来判断是显示登录界面还是处理登录逻辑。
当GET方式请求时显示登录界面，其他方式请求时则处理登录逻辑，如查询数据库、验证登录信息等。
*/