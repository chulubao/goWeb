package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	//t := template.Must(template.ParseFiles("html/layout.html","html/index.html"))
	/*包装在Must()中，它会自动在有err的时候panic，无错的时候只返回其中的*Template。*/
	/*ParseGlob:用于批量解析文件。*/
	t := template.Must(template.ParseGlob("html/*.html"))


	list :=[]string{"Go", "Python", "PHP", "JavaScript"}

	// 渲染模板，发送响应
	t.ExecuteTemplate(w, "layout", list)

	for i := 0; i < len(list); i++ {
		fmt.Printf(" value:%s\n", list[i])
	}



}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleAction)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/