package main
import ("github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	/*==========================[方法三]===============================*/

	/**
	 *[1] 默认配置解析
	 *beego 默认会解析当前应用下的 conf/app.conf 文件.
	 *[2] 通过这个文件(app.conf),你可以初始化很多 beego 的默认参数：
	 *viewspath = "info"  //info 为目录(文件夹)名
	 *httpport = 9090
	 */

	this.TplName = "info.html"
	/*================================================================*/



}
func (this *MainController)Post()  {
	this.Ctx.WriteString("hello chu zong !\n")
}

func main() {

	beego.Router("/login", &MainController{})
	beego.Run()
}

/**
*http://localhost:9090/login
*http://192.168.219.134:9090/login
*/

/*
[root@localhost handleHtmlRequests2]# go build main.go
[root@localhost handleHtmlRequests2]# ./main
[root@localhost handleHtmlRequests2]# go clean
*/