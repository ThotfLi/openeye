package dbops

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

func AddCommont(videoUUID string,userID uint,data string) error {
	uuid,err := uuid.NewV4()
	if err != nil {
		fmt.Println("[ERROR] uuid invition fail")
		return err
	}

	newCom := Comment{Video:&VideoInfo{VideoInfoId:videoUUID},
	                  Author:&User{Id:userID},
	                  Content:data,
	                  CommentId:uuid.String(),
					}

	if _,err := O.Insert(&newCom);err != nil {
		fmt.Println("[ERROR] Addcommont err")
		return err
	}
	return nil
}

func GetCommonts(videoUUID string)([]Comment,error){
	var comments []Comment
	qb,err := orm.NewQueryBuilder("mysql")
	if err != nil {
		fmt.Println("[ERROR] QueryBuilder Initialization fail")
		return nil,err
	}

	qb.Select("comment.Comment_id",
		"comment.content",
		"comment.create_time",
		"comment.author_id").
		From("comment").
		Where("video_id = ?").
		OrderBy("create_time")

	sql := qb.String()

	if _,err := O.Raw(sql,videoUUID).QueryRows(&comments);err != nil {
		fmt.Println("[ERROR] func GetCommonts Raw err")
		return nil,err
	}

	return comments,nil
}

func DelCommont(commontID string) error{
	com := Comment{CommentId:commontID}
	if _,err := O.Delete(&com); err != nil {
		fmt.Println("[ERROR] del comment fail")
		return err
	}
	return nil

}