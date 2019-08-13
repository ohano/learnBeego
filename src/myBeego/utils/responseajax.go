package utils

import "github.com/astaxie/beego"

type ResponseAjax struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data map[interface{}]interface{} `json:"data"`
}


func (r ResponseAjax) AjaxReturn(code , message string, data map[interface{}]interface{}) ResponseAjax {
	ra := ResponseAjax{Code: code, Message: message, Data: data}
	return ra
}

func DoResponse(c *beego.Controller, code, message string, data map[interface{}]interface{})  {
	ra := ResponseAjax{}.AjaxReturn(code, message, data)
	c.Data["json"] = ra
	c.ServeJSON()
}