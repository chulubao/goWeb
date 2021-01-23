package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)


func loginFunc(response http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method) //获取请求的方法
	if request.Method == "GET" {
		tmp ,err := template.New("login.html").ParseFiles("./html/login.html")
		t := template.Must(tmp,err)
		t.Execute(response, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		//不好用:fmt.Println("username:", request.Form["username"])
		fmt.Println("[1]username:", request.FormValue("username"))
		//不好用:fmt.Println("password:", request.Form["password"])
		fmt.Println("[1]password:", request.FormValue("password"))

		fmt.Println("[2]username:", template.HTMLEscapeString(request.Form.Get("username"))) //输出到服务器端
		fmt.Println("[2]password:", template.HTMLEscapeString(request.Form.Get("password"))) //输出到服务器端

		fmt.Println("[3]username:",template.HTMLEscaper(request.Form.Get("username"),"\n[3]password: ",request.Form.Get("password")))//输出到服务器端

		template.HTMLEscape(response, []byte(request.Form.Get("username"))) //输出到客户端
		template.HTMLEscape(response, []byte("\n")) //输出到客户端
		template.HTMLEscape(response, []byte(request.Form.Get("password"))) //输出到客户端

	}


}

func main() {
	http.HandleFunc("/login", loginFunc)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
http://localhost:9090/login
http://192.168.219.134:9090/login
*/

/*

Tips:
[2]Request本身也提供了FormValue()函数来获取用户提交的参数。
如r.Form["username"]也可写成r.FormValue("username")。
[2]调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。
[3]r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
*/

/*
[1]通过上面的代码我们可以看出获取请求方法是通过r.Method来完成的，这是个字符串类型的变量，返回GET, POST, PUT等method信息。
[2]login函数中我们根据r.Method来判断是显示登录界面还是处理登录逻辑。
当GET方式请求时显示登录界面，其他方式请求时则处理登录逻辑，如查询数据库、验证登录信息等。
*/

/*
4.3 预防跨站脚本

现在的网站包含大量的动态内容以提高用户体验，比过去要复杂得多。所谓动态内容，就是根据用户环境和需要，Web应用程序能够输出相应的内容。动态站点会受到一种名为“跨站脚本攻击”（Cross Site Scripting, 安全专家们通常将其缩写成 XSS）的威胁，而静态站点则完全不受其影响。

攻击者通常会在有漏洞的程序中插入JavaScript、VBScript、 ActiveX或Flash以欺骗用户。一旦得手，他们可以盗取用户帐户信息，修改用户设置，盗取/污染cookie和植入恶意广告等。

对XSS最佳的防护应该结合以下两种方法：一是验证所有输入数据，有效检测攻击(这个我们前面小节已经有过介绍);另一个是对所有输出数据进行适当的处理，以防止任何已成功注入的脚本在浏览器端运行。

那么Go里面是怎么做这个有效防护的呢？Go的html/template里面带有下面几个函数可以帮你转义

func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
func HTMLEscapeString(s string) string //转义s之后返回结果字符串
func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
*/