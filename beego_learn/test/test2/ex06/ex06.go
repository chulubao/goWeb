package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, request *http.Request) {
	// 解析url传递的参数
	request.ParseForm()
	for key, val := range request.Form {
		fmt.Println("key:", key)
		fmt.Println("val:", strings.Join(val, ""))
		fmt.Println("-------------------------------------")
	}


	fmt.Fprintf(w,"用户名:%v\n",request.FormValue("username"))
	fmt.Fprintf(w,"密码:%v\n",request.FormValue("password"))

	/*
	//方法二 输出到客户端
	name :=request.Form["username"]
	pass :=request.Form["password"]

	for i,v :=range name{
		fmt.Println(i)
		fmt.Fprintf(w,"用户名:%v\n",v)
	}
	for k,n :=range pass{
		fmt.Println(k)
		fmt.Fprintf(w,"密码:%v\n",n)
	}*/
}

func login(w http.ResponseWriter, request *http.Request) {

	fmt.Println("[1]method:", request.Method)

	if request.Method == "GET" {

		fmt.Println("[2]")

		t, _ := template.ParseFiles("./html/ex06.html")
		// func (t *Template) Execute(wr io.Writer, data interface{}) error {
		t.Execute(w, nil)
	} else {

		fmt.Println("[3]")

		request.ParseForm()
		fmt.Println("username:", request.Form["username"])
		fmt.Println("password:", request.Form["password"])
	}
}
func main() {
	
	http.HandleFunc("/login", login)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}


/**
*http://localhost:9090/login
*http://192.168.219.134:9090/login
*/