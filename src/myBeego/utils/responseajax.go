package utils

import "github.com/astaxie/beego"

type ResponseAjax struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data struct{} `json:"data"`
}


func (r ResponseAjax) AjaxReturn(code , message string, data struct{}) ResponseAjax {
	ra := ResponseAjax{Code: code, Message: message, Data: data}
	return ra
}

func DoResponse(c *beego.Controller, code, message string, data struct{})  {
	ra := ResponseAjax{}.AjaxReturn(code, message, data)
	c.Data["json"] = ra
	c.ServeJSON()
}