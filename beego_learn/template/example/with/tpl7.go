package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"html/template"
	"net/http"
	"log"
)


func sliceHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html"))

	/*使用 make创建空的字符串切片*/
	dayWeek := make([]string, 0)

	dayWeek = append(dayWeek, "Fri")
	dayWeek = append(dayWeek, "Sat")
	dayWeek = append(dayWeek, "Sun")
	dayWeek = append(dayWeek, "")
	t.ExecuteTemplate(w, "layout", dayWeek)
}

func main() {
	// 设置 处理函数
	http.HandleFunc("/", sliceHandler)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*
http://localhost:8081
curl http://localhost:8081
http://192.168.147.186:8081
curl  http://192.168.147.186:8081
*/