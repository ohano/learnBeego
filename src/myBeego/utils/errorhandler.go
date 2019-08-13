package utils

import (
	"github.com/astaxie/beego"
)

type ErrorHandle struct {
	beego.Controller
}

func (e ErrorHandle) ErrorPage() {
	e.TplName = "error.tpl"
	_ = e.Render()
}