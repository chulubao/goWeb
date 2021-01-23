package main

import (
	"github.com/astaxie/beego"
	"strconv"
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
/*this.Ctx.WriteString("hello chu zong !\n")*/
/*GetString里面的参数和我们刚刚表单设置的属性 一模一样（注意：名字必须一模一样，为了避免坑大小写都要相同）*/
	Name:=this.GetString("Name")
	Pwd:=this.GetString("Pwd")
	Sex:=this.GetString("Sex")
	Age,err:=this.GetInt("Age")
	if err==nil {
		this.Ctx.WriteString("Name="+Name+"\nPwd="+Pwd+"\nSex="+Sex+"\nAge="+strconv.Itoa(Age))
	}else{
		this.Ctx.WriteString("Name="+Name+"\nPwd="+Pwd+"\nSex="+Sex)
	}

}

func main() {
	/*==========================[方法二]===============================*/
	/*自定义模板路径的方法*/
	beego.BConfig.WebConfig.ViewsPath="html"
	/*================================================================*/
	beego.Router("/register", &MainController{})
	beego.Run()
}

/**
*http://localhost:8080/register
 */