package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"html/template"
	"net/http"
	"log"
)

func arrayHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Ths", "Fri", "Sat", "Sun"}
	t.ExecuteTemplate(w, "layout", daysOfWeek)
}

func sliceHandler1(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html"))

	/*使用切片字面量创建空的字符串切片*/
	dayWeek := []string{}

	dayWeek = append(dayWeek, "Mon")
	dayWeek = append(dayWeek, "Tue")
	dayWeek = append(dayWeek, "Wed")

	t.ExecuteTemplate(w, "layout", dayWeek)
}

func sliceHandler2(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/layout.html"))

	/*使用 make创建空的字符串切片*/
	dayWeek := make([]string, 0)

	dayWeek = append(dayWeek, "Fri")
	dayWeek = append(dayWeek, "Sat")
	dayWeek = append(dayWeek, "Sun")
	t.ExecuteTemplate(w, "layout", dayWeek)
}

func main() {
	// 设置 处理函数
	http.HandleFunc("/", sliceHandler2)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/