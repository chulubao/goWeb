package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)



func handleMapAction(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("html/index.html")
	// 设置模板数据
	data := map[string]interface{}{
		"User": "小韩说课",
		"List": []string{"Go", "Python", "PHP", "JavaScript"},
	}
	// 渲染模板，发送响应
	t.Execute(w, data)
}

/*********************************************************
go提供了ExecuteTemplate方法，用于执行指定名字的模板。
例如加载layout.html模板的时候，可以指定layout.html
======================================================
在模板文件中，使用了define这个action给模板文件命名了。
虽然我们ParseFiles方法返回的模板对象t的名字还是layout.html，
但是ExecuteTemplate执行的模板却是html文件中定义的layout。

**********************************************************/
func handleMapAction1(w http.ResponseWriter, r *http.Request) {
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


func main_tpl1() {
	// 设置 处理函数
	http.HandleFunc("/", handleMapAction1)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/