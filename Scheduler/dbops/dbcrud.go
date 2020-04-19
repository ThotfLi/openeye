package dbops

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"openeye/api/dbops"
)

func GetWaitDelUsers() ([]WaitDelUser,error) {
	var videos []WaitDelUser
	qb,err := orm.NewQueryBuilder("mysql")
	if err != nil {
		fmt.Println("[ERROR] VideoDeleteProducer.NewQueryBuilder is err",err)
		return nil, err
	}

	qb.Select("wait_del_user.video_id").From("wait_del_user").Limit(3)
	sql := qb.String()
	_,err = dbops.O.Raw(sql).QueryRows(&videos)
	if err != nil {
		fmt.Println("[ERROR] videoDeleteProducer.QueryRows is err",err)
		return nil, err
	}
	return videos,nil
}


func DeleteWaitDelUser(d string)error{
	var newWait WaitDelUser
	newWait.VideoID = d

	_,err := dbops.O.Delete(&newWait)
	if err != nil {
		return err
	}
	return nil
}


func InsertWaitDelUser(id string) error {
	var newTable WaitDelUser
	newTable.VideoID = id

	_,err := O.Insert(&newTable)
	if err != nil {
		return err
	}
}
