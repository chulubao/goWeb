package main
import ("github.com/astaxie/beego")
type MainController struct {
	beego.Controller
}

func (cntl *MainController) Get() {
	cntl.Ctx.WriteString("hello world!\n" + "this is test beego!")
}

func main() {
	beego.Router("/hello", &MainController{})
	beego.Run()
}
/**
*http://localhost:8080/hello
*http://192.168.219.134:8080/hello
*/