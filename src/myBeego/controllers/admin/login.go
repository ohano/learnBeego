package admin

import (
	"fmt"
	"myBeego/controllers"
	"myBeego/models/admin"
	"myBeego/utils"
)

type LoginController struct {
	controllers.AdminBaseController
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
	data struct{}
)

func (l *LoginController) PostLogin() {
	l.XSRFToken()
	xsrfToken := l.CheckXSRFCookie()
	fmt.Println(xsrfToken)
	if !xsrfToken {
		utils.DoResponse(&l.Controller, code, "token错误", data)
		return
	}
	username := l.GetString("username")
	password := l.GetString("password")
	managerInfo, err := admin.GetManagerByUsername(username)

	fmt.Println(err)
	if err != nil {
		utils.DoResponse(&l.Controller, code, "用户名或密码错误", data)
		return
	}
	sig := utils.String2Md5(utils.String2Md5(password + managerInfo.Salt))
	if sig != managerInfo.Password {
		utils.DoResponse(&l.Controller, code, "用户名或密码错误", data)
		return
	}
	code = "1000"
	utils.DoResponse(&l.Controller, code, "登录成功", data)
	l.SetSession("manager", managerInfo)
	return
}
