package main


/*模板引擎（template engine）: 将数据嵌入到 HTML 中 */
import (
	"net/http"
	"encoding/xml"
	"github.com/unrolled/render"  // or "gopkg.in/unrolled/render.v1"
)

type ExampleXml struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}


/* go  get   github.com/unrolled/render */

func main() {
	r := render.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w, http.StatusOK, "Plain text here")
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/jsonp", func(w http.ResponseWriter, req *http.Request) {
		r.JSONP(w, http.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
	})

	mux.HandleFunc("/xml", func(w http.ResponseWriter, req *http.Request) {
		r.XML(w, http.StatusOK, ExampleXml{One: "hello", Two: "xml"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./templates called "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HTML(w, http.StatusOK, "example", "World")
	})

	//// 开启监听（监听浏览器请求）
	http.ListenAndServe(":8081", mux)
}


/*
http://localhost:8081
curl http://localhost:8081
http://192.168.147.186:8081
curl  http://192.168.147.186:8081
*/