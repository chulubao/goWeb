package main
import (
"github.com/astaxie/beego"
)



type User struct {
	Name string  //首字母一定要大写
	Age  int
	Sex  string
}

type HelloControllers struct {
	beego.Controller
}

func (hello *HelloControllers) Get() {
	user:= &User{"zhang san",30,"body"}
	/*
	hello.Data["json"] =user
	//调用 ServeJSON 之后，会设置 content-type 为 application/json，然后同时把数据进行 JSON 序列化输出.
	hello.ServeJSON()
	// {{.json}}
	*/

	
	hello.Data["xml"] =user
	//调用 ServeXML 之后，会设置 content-type 为 application/xml，同时数据会进行 XML 序列化输出.
	hello.ServeXML()
	//{{.xml}}


	/*
	hello.Data["jsonp"] = user
	hello.ServeJSONP()
	*/

	//模板路径，默认值是 ./views/index.tpl
	hello.TplName = "index.tpl"
}


func main() {
	 /*beego.Router("/", &MainController{})*/
	// "/helloWorld" 是我们访问的路径,比如说我们想要调用HelloControllers这个控制器
	//需要在浏览器输入http://127.0.0.1:8090/helloWorld
	beego.Router("/helloWorld", &HelloControllers{})
	beego.Run()
}

/********************************************************************************
http://127.0.0.1:8080/helloWorld
http://localhost:8080/helloWorld
**********************************************************************************
结束端口号占用:
[root@localhost ~]# lsof -i:8080
COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
beegoHelloWorld 18628 root    3u  IPv6  90885      0t0  TCP *:8090 (LISTEN)
[root@localhost ~]# kill 18628
***********************************************************************************
[root@localhost ~]# bee new <项目名称>;
我输入
[root@localhost ~]# bee new  beegoHelloWorld
bee工具会自动在$GOPATH/src目录下生成beegoHelloWorld项目(目录结构完整)
*/