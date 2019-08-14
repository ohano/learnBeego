package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "myBeego/models"
	_ "myBeego/routers"
	"myBeego/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.ErrorController(&utils.ErrorHandle{})
	beego.Run()
}

