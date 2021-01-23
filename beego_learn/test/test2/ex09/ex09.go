package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func Add(left int, right int) int {
	return  left+right;

}

func tmplFunc(response http.ResponseWriter, request *http.Request) {
	funcMap := template.FuncMap{"Add":Add}
	t := template.New("ex09.html").Funcs(funcMap)
	tmplt, err := t.ParseFiles("./html/ex09.html")
	if err!=nil {
		fmt.Fprintln(response, err)
		return
	}

	tmplt.Execute(response, nil)
}


func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/html", tmplFunc)
	server.ListenAndServe()
}


/**
*http://localhost:8080/html
*http://192.168.219.134:8080/html
*/

//[1]
//【模板标签】
//模板标签用"{{"和"}}"括起来
//------------------------------------
// [2]
//【注释】
// {{/* a comment */}}
// 使用“{{/*”和“*/}}”来包含注释内容
//------------------------------------
//[3]
/*
【模板函数】 或叫 【管道函数】
用法1：
{{FuncName1}}
此标签将调用名称为“FuncName1”的模板函数（等同于执行“FuncName1()”，不传递任何参数）并输出其返回值。
用法2：
{{FuncName1 "参数值1" "参数值2"}}
此标签将调用“FuncName1("参数值1", "参数值2")”，并输出其返回值
用法3：
{{.Admpub|FuncName1}}
此标签将调用名称为“FuncName1”的模板函数（等同于执行“FuncName1(this.Admpub)”，将竖线“|”左边的“.Admpub”变量值作为函数参数传送）并输出其返回值。
 */