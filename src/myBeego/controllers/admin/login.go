package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) GetLogin()  {
	l.Data["username"] = "admin"
	l.Data["password"] = "xm123123"
	xsrfToken := l.XSRFToken()
	l.Data["token"] = xsrfToken
	l.TplName = "admin/login.tpl"
}

func (l *LoginController) PostLogin() {
	l.Data["username"] = l.GetString("username")
	l.Data["password"] = l.Ctx.Request.Form.Get("password")
	l.Ctx.WriteString(l.Data["username"].(string))
	l.Ctx.WriteString(l.Data["password"].(string))
	o := orm.NewOrm()
	fmt.Println(l.Data)
}
