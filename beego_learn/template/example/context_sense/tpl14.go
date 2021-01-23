package main

import (
	"html/template"
	"net/http"
)

/*
上面test.html中有4个不同的环境，分别是html环境、url的path环境、url的query环境以及js环境。
虽然对象都是{{.}}，但解析执行后的值是不一样的。如果使用curl获取结果(curl http://localhost:8081)
========================================================================================
上下文感知的自动转义能让程序更加安全

*/
func handleAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/test.html")
	if err != nil {
		panic(err)
	}

	content := `I need: <i>"What's up?"</i>`

    /*未转义*/
	//t.Execute(w, content)

	/*转义*/
	t.Execute(w, template.HTML(r.FormValue(content)))

}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleAction)
	//// 开启监听（监听浏览器请求）
	http.ListenAndServe(":8081", nil)
}

/*curl http://localhost:8081*/