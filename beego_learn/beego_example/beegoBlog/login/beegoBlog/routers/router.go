package routers

import (
	"beegoBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/default", &controllers.MainController{})   /*http://localhost:8080/default*/
	beego.Router("/login", &controllers.LoginController{})    /*http://localhost:8080/login*/
}
