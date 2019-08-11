package routers

import (
	"github.com/astaxie/beego"
	"myBeego/controllers"
	"myBeego/controllers/admin"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
}
