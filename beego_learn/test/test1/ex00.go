package main

import "github.com/astaxie/beego"


type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main1() {
	beego.Router("/", &MainController{})
	beego.Run()
}
/**
*http://localhost:8080/
*http://192.168.219.134:8080/
*/