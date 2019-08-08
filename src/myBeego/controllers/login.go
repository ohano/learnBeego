package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) GetLogin()  {
	l.Data["username"] = "admin"
	l.Data["password"] = "xm123123"
	l.TplName = "login.tpl"
}

func (l *LoginController) PostLogin() {
	l.Data["user_name"] = l.Ctx.Request.Form.Get("username")
	l.Data["password"] = l.Ctx.Request.Form.Get("password")
	fmt.Println(l.Data)
}
