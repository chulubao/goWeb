package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"html/template"
	"net/http"
	"log"
)


func sliceHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html", "html/index.html","html/content.html"))

	dayWeek := []string{"Mon", "Tue", "Wed", "Ths", "Fri", "Sat", "Sun"}

	dayWeek = append(dayWeek, "end")


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