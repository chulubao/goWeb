package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

func handleMapAction2(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("html/layout.html")

	// 设置模板数据
	data := map[string]interface{}{
		"User": "小韩说课",
		"List": []string{"Go", "Python", "PHP", "JavaScript"},
	}

	fmt.Println(t.Name())

	// 渲染模板，发送响应
	t.ExecuteTemplate(w, "layout", data)
}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleMapAction2)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/