package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"html/template"
	"net/http"
	"log"
	"math/rand"
	"time"
)

func tmpHandler(w http.ResponseWriter, r *http.Request){
	t, _ :=template.ParseFiles("html/layout.html")
	rand.Seed(time.Now().Unix())

	var isOk bool
	isOk=rand.Intn(10) > 5

	/*go提供了ExecuteTemplate方法，用于执行指定名字的模板。例如加载layout.html模板的时候，可以指定layout.html*/
	t.ExecuteTemplate(w, "layout", isOk)

	/*休眠2秒*/
	/*time.Sleep(time.Duration(2)*time.Second)*/

}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", tmpHandler)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*http://localhost:8081*/