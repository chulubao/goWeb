package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	t := template.New("test")
	if rand.Intn(10) > 5 {
		t , _ = template.ParseFiles("html/content.html", "html/red.html")
	} else {
		t, _ = template.ParseFiles("html/content.html")
	}

	t.Execute(w,"")

}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleAction)
	//// 开启监听（监听浏览器请求）
	http.ListenAndServe(":8081", nil)
}

/*http://localhost:8081*/