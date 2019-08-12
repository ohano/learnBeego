package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbDriver := beego.AppConfig.String("DB_DRIVER")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPass := beego.AppConfig.String("DB_PASS")
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbCharset := beego.AppConfig.String("DB_CHARSET")
	dbname := beego.AppConfig.String("DB_DATABASE")
	dsn := dbUser + ":" + dbPass + "@tcp("+ dbHost +":"+ dbPort +")/"+ dbname +"?charset=" + dbCharset
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//_ = orm.RegisterDriver(dbDriver, orm.DRMySQL)
	err := orm.RegisterDataBase("default", dbDriver, dsn)
	if err != nil {
		fmt.Println(err)
	}
	orm.RegisterModel()
}
