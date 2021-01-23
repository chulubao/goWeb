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

/****************************************************************************
[1]Tag给属性设置别名,它可以让结构体属性对应到表单中的name属性.
比如<input type="text" name="college" value="">  name属性值college首字母是小写,
但是结构体的属性University必须是大写,不然就解析不到.
如果在不更改表单name属性值college的情况下，就要用到form来对应,
form后面的字段必须要和对应表单中的college大小写完全一样.
如: University string `form:"college" `
-----------------------------------------------------------------------------
[2]不想解析表单某个属性,我们可以使用结构体Tag 在属性后添加`form:"-"` ;还有一种方法就是开头字母小写.
-----------------------------------------------------------------------------
[3]beego提供了ParseForm方法用来解析表单数据,调用ParseForm 这个方法的时候,传入的参数必须为一个 struct 的指针，
否则对 struct 的赋值不会成功并返回xx must be a struct pointer的错误.
*****************************************************************************/
type UserInfo struct {  /*Tag*/
	Name string		"用户名"
	Pwd  string		"用户密码"
	Sex  string		"用户性别"
	Age int			"用户年龄"
	University string `form:"college" `
	Work string	` form:"-" `                /*不想解析此属性*/
	idNumber  string                        /*不想解析此属性*/
	Level     string	"职员"
}

func (this *MainController)Post()  {
	/*this.Ctx.WriteString("hello chu zong !\n")*/
	var userMsg UserInfo
	error:=this.ParseForm(&userMsg)
     if error!=nil{
     	this.Ctx.WriteString("出错了！")
     }else{
     	this.Ctx.WriteString("我是结构体\n")
     	this.Ctx.WriteString("Name="+userMsg.Name+
     		"\nPwd="+userMsg.Pwd+
     		"\nSex="+userMsg.Sex+
     		"\nAge="+strconv.Itoa(userMsg.Age)+
     		"\nCollege="+userMsg.University+
			"\nWork="+userMsg.Work+          /*不想解析此属性*/
			"\nIdNumber="+userMsg.idNumber+  /*不想解析此属性*/
			"\nLevel="+userMsg.Level )
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

/*************************************
*http://localhost:8080/register
*----------------结果-----------------*
我是结构体
Name=clj
Pwd=123456
Sex=woman
Age=30
College=heidaxue
Work=                     //不想解析此属性
IdNumber=                 //不想解析此属性
Level=zhiyuan
****************************************/