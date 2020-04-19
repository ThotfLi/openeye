package dbops

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
//数据库连接初始化

var O orm.Ormer

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:1@tcp(127.0.0.1:3306)/openeye?charset=utf8&loc=Local", 30)

	// register model
	orm.RegisterModel(new(User),new(VideoInfo),new(Comment),new(Sessions))

	// create table
	orm.RunSyncdb("default",false, true)

	O = orm.NewOrm()
}
