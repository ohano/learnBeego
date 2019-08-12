package admin

import (
	"fmt"
	"github.com/astaxie/beego"
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
	orm.RegisterModelWithPrefix(beego.AppConfig.String("DB_PREFIX"), new(Manager))
}

func (m *Manager) TableName() string {
	return "manager"
}

func GetManagerById (id int) (manager *Manager, err error) {
	o := orm.NewOrm()
	manager = &Manager{Id: id}
	//manager := new(Manager)
	err = o.Read(manager)
	//qs := o.QueryTable(new(Manager))
	//qs := o.QueryTable(new(Manager))
	//err = qs.Filter("id", id).One(manager)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return manager, err
}
