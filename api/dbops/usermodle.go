package dbops

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

//添加一个用户
func AddUserCreate(loginName string,pwd string)error{
	newUser := User{LoginName:loginName,Pwd:pwd}
	_,err := O.Insert(&newUser)
	if err != nil {
		return err
	}
	return nil
}

//获取一个用户密码
func GetUser(loginName string)(string,error){
	user := User{}
	qs := O.QueryTable("user")

	err := qs.Filter("login_name",loginName).One(&user)
	if err != nil {
		return "",err
	}

	return user.Pwd,nil
}

//删除一个用户
func DeleteUser(loginName string,pwd string)error{
	qs := O.QueryTable("user")

	cnt, err := qs.Filter("login_name",loginName).Filter("pwd",pwd).Delete()
	if err != nil {
		return err
	}
	fmt.Println(cnt)
	return nil
}