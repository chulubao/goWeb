package main

import (
	"html/template"
	"net/http"
)
/**
(1)go 统一使用了 {{ 和 }} 作为左右标签，没有其他的标签符号。
(2)使用 . 来访问当前位置的上下文。
(3)使用 $ 来引用当前模板根级的上下文。
(4)使用 $var 来访问创建的变量。
(5){{.}} 表示当前字段。
(6){{.Name}} 表示某个结构体的 Name 字段,结构体字段需能被外部访问,即字段首字母大写。
(7)
func New(name string) *Template
创建一个名为name 的模板。
(8)
func (t Template) Parse(text string) (Template, error)
将字符串text解析为模板。
text里面需要填充的字段要严格按照template包规定的那样填写。
如：fsdnfdsfd{{.Name}}fdsfkksdsdfdfsds{{if .Age}}fgsgdgdf{{else}}fdsfdsfjids{{end}}
Name,Age就是需要被填充的字段
(9)
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
Execute方法将解析好的模板应用到data上，并将输出写入wr。
如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行
*/
func SayHello(response http.ResponseWriter, request *http.Request) {
	name := "chulujian"
	tmpl, _ := template.New("tmpName").Parse("大家好，我是{{.}}")
	tmpl.Execute(response, name)
}

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8080", nil)
}

/**
*http://localhost:8080
*http://192.168.219.134:8080
*/