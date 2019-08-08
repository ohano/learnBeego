package routers

import (
	"github.com/astaxie/beego"
	"myBeego/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{}, "get:GetLogin;post:PostLogin")
}
