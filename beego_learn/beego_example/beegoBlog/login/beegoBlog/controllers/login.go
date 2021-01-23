package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "Login.html"
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		//c.Ctx.SetCookie("uname", "", -1, "/")
		//c.Ctx.SetCookie("pwd", "", -1, "/")
		//c.Redirect("/default", 301)  /*重定向到index.tpl*/
		c.Ctx.WriteString("退出登录成功！")
		return
	}
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autologin := c.Input().Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {

		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		//c.Ctx.SetCookie("uname", uname, maxAge, "/")
		//c.Ctx.SetCookie("pwd", pwd, maxAge, "/")

		c.Data["UserName"]=uname
		c.Data["Password"]=pwd
		c.Data["MaxAge"]=maxAge
		//c.Redirect("/default", 301)  /*重定向到最初的index.tpl(参数不会传入)*/
		c.TplName = "index.tpl"  /*传入参数重新加载index.tpl*/
	}else {
		c.Ctx.WriteString("登录用户名或登录密码错误!")
	}

	return
}
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {

		return true
	}

	return false
}
