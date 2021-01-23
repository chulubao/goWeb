package main
import ("github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//this.Ctx.WriteString("hello world!\n")
	/*================================[方法一]================================*/
	/**
	 * beego的默认参数
	 *ViewsPath
	 *模板路径，默认值是 views
	 */

	/**
	 *配置默认目录:  ./views/login.html
	 *无views目录（文件夹）,需要手动创建.
	 */
	this.TplName = "login.html"
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
*http://localhost:8080/login
*http://192.168.219.134:8080/login
*/

/*
[root@localhost handleHtmlRequests0]# go build main.go
[root@localhost handleHtmlRequests0]# ./main
[root@localhost handleHtmlRequests0]# go clean
*/