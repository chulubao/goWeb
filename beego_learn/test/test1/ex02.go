package main
import ("github.com/astaxie/beego")
type MainController struct {
	beego.Controller
}
func (cntl *MainController) GetData() {
	cntl.Ctx.WriteString("id=")
	id := cntl.GetString("id")
	cntl.Ctx.WriteString(id)

	cntl.Ctx.WriteString("\n")

	cntl.Ctx.WriteString("name=")
	name := cntl.Input().Get("name")
	cntl.Ctx.WriteString(name)
}


func main() {
	beego.Router("/getdata", &MainController{}, "get:GetData")
	beego.Run()
}
/**
*http://localhost:8080/getdata?id=clj&name=chulujian
*http://192.168.219.134:8080/getdata?id=clj&name=chulujian
*/