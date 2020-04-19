package dbops

import "github.com/astaxie/beego/orm"

var O orm.Ormer

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:1@tcp(127.0.0.1:3306)/openeye?charset=utf8&loc=Local", 30)

	// register model
	orm.RegisterModel(new(WaitDelUser))

	// create table
	orm.RunSyncdb("default",true, true)

	O = orm.NewOrm()
}