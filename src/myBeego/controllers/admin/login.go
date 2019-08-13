package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"myBeego/models/admin"
	"myBeego/utils"
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
	_ = l.Render()
}

func (l *LoginController) PostLogin() {
	username := l.GetString("username")
	password := l.Ctx.Request.Form.Get("password")
	managerInfo, err := admin.GetManagerByUsername(username)
	if err != nil {
		l.Abort("Page")
	}
	sig := utils.String2Md5(utils.String2Md5(password + managerInfo.Salt))
	if sig != managerInfo.Password {
		data := utils.ResponseAjax{}.AjaxReturn("1001", "密码错误", "")
		fmt.Println(data)
		l.Data["json"] = data
		l.ServeJSON()
	}
	l.Data["json"] = utils.ResponseAjax{}.AjaxReturn("1000", "登录成功", "")
	l.ServeJSON()
}
