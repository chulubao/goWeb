package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type Info struct {
	InfoTitle  string
	InfoContent string
	InfoTimes  int
}

func tempFunc(response http.ResponseWriter, request *http.Request) {
	tmplt, err := template.ParseFiles("./html/ex07.html")
	if err!=nil {
		fmt.Fprintln(response, err)
		return
	}

    infStu:=Info{InfoTitle:"模板",InfoContent:"传参",InfoTimes:12}
    
	tmplt.Execute(response, infStu)
}


func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/html", tempFunc)
	server.ListenAndServe()
}


/**
*http://localhost:8080/html
*http://192.168.219.134:8080/html
*/