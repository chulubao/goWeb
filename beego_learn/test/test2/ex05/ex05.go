package main

import (
	"net/http"
	"html/template"
)

func tmpfunc(response http.ResponseWriter, request *http.Request) {
	/**
	 *(1)调用ParseFiles函数解析模板文件时，Go会创建一个新的模板，并将给定的模板文件的名字作为新模板的名字;
	 *如果该函数中传入了多个文件名，那么也只会返回一个模板，而且以第一个文件的文件名作为模板的名字;
	 *至于其他文件对应的模板则会被放到一个map中。
	 *(2)
	 * t, _ := template.ParseFiles("hello.html")
	 *等价于
	 *t := template.New("hello.html")
     *t, _ = t.ParseFiles("hello.html")
	 *(3)文件名一致(在New  ParseFiles ExecuteTemplate 函数中)
	 *tm,_:= template.New("demo.html").ParseFiles("./html/ex04.html", "./html/demo.html")
	 *tm.ExecuteTemplate(response, "demo.html","*at demo show- :hello chulujian!")
	 *(4)执行模板
	 *[1]func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
     *Execute方法将解析好的模板应用到data上，并将输出写入wr。
	 *如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行.
	 *[2]
	 *func (t Template) ExecuteTemplate(wrio.Writer, namestring, data interface{})error
	 *ExecuteTemplate方法类似Execute，但是使用名为name的t关联的模板产生输出。
	 */

	/*//[一]
	t, err := template.ParseFiles("./html/ex05.html")
	if err!=nil {
		fmt.Fprintln(response, err)
		return
	}
	t.Execute(response, "hello chulujian!!")
	*/




	/*//[二]
	tmp, _ := template.ParseFiles("./html/ex04.html", "./html/ex05.html")
	tmp.ExecuteTemplate(response, "ex05.html","at ex05 show :hello chulujian!")
	*/

	/*//[三]
	tmp := template.New("ex05.html")
	tt, _ := tmp.ParseFiles("./html/ex04.html", "./html/ex05.html")
	tt.ExecuteTemplate(response, "ex05.html","-at ex05 show- :hello chulujian!")
	*/

    //[四]
	ttt,_:= template.New("ex05.html").ParseFiles("./html/ex04.html", "./html/ex05.html")
	ttt.ExecuteTemplate(response, "ex05.html","*at ex05 show- :hello chulujian!")
}


func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/html", tmpfunc)
	server.ListenAndServe()
}


/**
*http://localhost:8080/html
*http://192.168.219.134:8080/html
*/