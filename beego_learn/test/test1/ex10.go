package main
import ("github.com/astaxie/beego")

type MainController struct {
	beego.Controller
}

func (cntl *MainController) Get() {
	cntl.Ctx.WriteString("hello world!\n")
	cntl.Ctx.WriteString(cntl.Ctx.Input.Param(":username"))
}

func main() {
	//beego.Router("/hello/?:id", &MainController{})
	beego.Router("/hello/?:username", &MainController{})
	beego.Run()
}

/**
*http://localhost:8080/hello/11
*http://192.168.219.134:8080/hello/11
*http://localhost:8080/hello/chubao
*/