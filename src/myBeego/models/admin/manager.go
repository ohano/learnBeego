package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Manager struct {
	Id	int
	Username string
	Password string
	Email string
	Salt string
}

func init() {
	orm.RegisterModel(new(Manager))
}

func GetManagerById (id int) (manager *Manager, err error) {
	o := orm.NewOrm()
	//manager = &Manager{Id : id}
	//err = o.Read(&manager)
	qs := o.QueryTable("xm_manager")
	err = qs.Filter("id", id).One(&manager)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return manager, nil
}
