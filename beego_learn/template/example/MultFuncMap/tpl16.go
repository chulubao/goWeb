package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"time"
)

func formatNow(format string) string {
	return time.Now().Format(format)
}


func addFunc(arg1 int ,arg2 int ) int {
	return (arg1+arg2)
}


func subFunc(arg1 int ,arg2 int ) int {
	return (arg1-arg2)
}

func multFunc(arg1 int ,arg2 int ) int {
	return (arg1*arg2)
}


/*
[1]函数的注入，必须要在parseFiles之前，因为解析模板的时候，需要先把函数编译注入。
[2]创建一个FuncMap类型的map，key是模板函数的名字，value是其函数的定义。
*/
func handleAction(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"showTime": formatNow, "sum":addFunc, "sub":subFunc, "mult":multFunc}

	/*将 FuncMap注入到模板中*/
	t := template.New("index.html").Funcs(funcMap)


	t =template.Must(t.ParseFiles("html/index.html"))


	/*函数传入参数:2006-01-02 15:04:05*/
	err := t.ExecuteTemplate(w, "index", "2006-01-02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}


}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleAction)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/