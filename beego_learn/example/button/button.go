package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)



func loginFunc(respose http.ResponseWriter, reques *http.Request) {
	fmt.Println("method:", reques.Method) //获取请求的方法
	if (reques.Method == "GET") {
		tmp ,err := template.New("button.html").ParseFiles("./button.html")
		tpl := template.Must(tmp,err)
		tpl.Execute(respose, nil)
	} else  if (reques.Method == "POST"){
		button := reques.FormValue("id_clicked")
		selectv :=reques.PostFormValue("sel")

		fmt.Println("button:"+button)
		fmt.Println("select:"+selectv)
		if (selectv == "idv1") {
			respose.Write([]byte("1AA!"))
		}

		if (selectv == "idv2") {
			respose.Write([]byte("2BB!"))
		}

		if (selectv == "idv3") {
			respose.Write([]byte("3CC!"))
		}

	}


}

func main() {
	http.HandleFunc("/btn", loginFunc)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
http://localhost:9090/btn
http://192.168.219.134:9090/btn
*/