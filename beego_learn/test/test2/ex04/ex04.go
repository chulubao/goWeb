package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func templateFunc(response http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./html/ex04.html")
	if err!=nil {
		fmt.Fprintln(response, err)
		return
	}
	t.Execute(response, "hello chulujian!!")
}


func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/html", templateFunc)
	server.ListenAndServe()
}


/**
*http://localhost:8080/html
*http://192.168.219.134:8080/html
*/