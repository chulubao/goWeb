package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type Msg struct {
	MsgTitle  string
	MsgContent string
	MsgTimes  int
}

func tmpFunc(response http.ResponseWriter, request *http.Request) {
	tmplt, err := template.ParseFiles("./html/ex08.html")
	if err!=nil {
		fmt.Fprintln(response, err)
		return
	}

	msgStu:=Msg{MsgTitle:"模板",MsgContent:"传参",MsgTimes:12}

	tmplt.Execute(response, msgStu)
}


func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/html", tmpFunc)
	server.ListenAndServe()
}


/**
*http://localhost:8080/html
*http://192.168.219.134:8080/html
*/

/*[1] 在模板中定义变量：变量名称用字母和数字组成，并带上“$”前缀，采用符号“:=”进行赋值。比如："{{$x := "OK"}} 或 {{$x := pipeline}}。
 *[2] {{$admpub}} 此标签用于输出在模板中定义的名称为“admpub”的变量。当$admpub本身是一个Struct对象时，可访问其字段。
*/