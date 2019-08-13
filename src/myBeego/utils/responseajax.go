package utils

type ResponseAjax struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func (r ResponseAjax) AjaxReturn(code , message string, data interface{}) ResponseAjax {
	ra := ResponseAjax{Code: code, Message: message, Data: data}
	return ra
}
