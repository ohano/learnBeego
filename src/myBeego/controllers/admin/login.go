package admin

import (
	"github.com/astaxie/beego"
	"myBeego/models/admin"
	"myBeego/utils"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) GetLogin()  {
	xsrfToken := l.XSRFToken()
	l.Data["token"] = xsrfToken
	l.TplName = "admin/login.tpl"
	_ = l.Render()
}
/**
登录
 */
var (
	code string = "1001"
	message string = ""
	data map[interface{}]interface{} = nil
)

func (l *LoginController) PostLogin() {
	//errToken := l.CheckXSRFCookie()
	//if !errToken {
	//	utils.DoResponse(&l.Controller, code, "token错误", data)
	//}
	username := l.GetString("username")
	password := l.GetString("password")
	managerInfo, err := admin.GetManagerByUsername(username)
	if err != nil {
		l.Abort("Page")
	}
	sig := utils.String2Md5(utils.String2Md5(password + managerInfo.Salt))
	if sig != managerInfo.Password {
		utils.DoResponse(&l.Controller, code, "密码错误", data)
	}
	code = "1000"
	utils.DoResponse(&l.Controller, code, "登录成功", data)
}
