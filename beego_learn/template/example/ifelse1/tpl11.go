package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	 t ,err := template.ParseFiles("html/layout.html")
	if err != nil {
		panic(err)
	}

	/*包装在Must()中，它会自动在有err的时候panic，无错的时候只返回其中的*Template。*/
	/*ParseGlob:用于批量解析文件。*/
	//t := template.Must(template.ParseGlob("html/*.html"))


	list :=[]string{"星期一",
		"星期二",
		"星期三",
		"星期四",
		"星期五",
		"星期六",
		"星期日"}

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

/*
Go标准库：Go template用法详解:https://www.cnblogs.com/f-ck-need-u/p/10053124.html
*/
/*http://localhost:8081*/