package controllers

import (
	"github.com/astaxie/beego"
	"time"
)


type AdminBaseController struct {
	beego.Controller
	isLogin bool
}

func (b *AdminBaseController) Prepare()  {
	b.Data["PageStartTime"] = time.Now()
	b.Data["AppName"] = "hanobeego"
	//TODO:验证权限，菜单之类
}
