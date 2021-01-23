package main
import ("github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	/*==========================[方法二]===============================*/
	/*
	[1] beego.BConfig.WebConfig.ViewsPath="tmp" ; 无this.TplName="path" ,   默认模板查找路径为： tmp/maincontroller/get.tpl
	[2]beego.BConfig.WebConfig.ViewsPath="tmp" ; this.TplName="login.html" ,默认模板查找路径为： tmp/login.html
	*/

	/*-----------------------------------------------------------------*/
          //(1)
	  this.TplName = "login.html"

	 //(2)注释掉this.TplName = "login.html" ; 默认模板查找路径为： tmp/maincontroller/get.tpl

	/*================================================================*/

}

func (this *MainController)Post()  {
	this.Ctx.WriteString("hello chu zong !\n")
}

func main() {

	/*==========================[方法二]===============================*/
	/*自定义模板路径的方法*/

	 beego.BConfig.WebConfig.ViewsPath="tmp"
	/*================================================================*/

	beego.Router("/login", &MainController{})
	beego.Run()
}

/**
*http://localhost:8080/login
*http://192.168.219.134:8080/login
*/

/*
[root@localhost handleHtmlRequests1]# go build main.go
[root@localhost handleHtmlRequests1]# ./main
[root@localhost handleHtmlRequests1]# go clean
*/
