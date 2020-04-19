package taskrunner

import (
	"openeye/Scheduler/dbops"
	"os"
	"strings"
	"sync"
)

func VideoDeleteProducer(da DataChan) error {
	//从待del_video 中拿到id
	//var videos []dbops.WaitDelUser
	//qb,err := orm.NewQueryBuilder("mysql")
	//if err != nil {
	//	fmt.Println("[ERROR] VideoDeleteProducer.NewQueryBuilder is err",err)
	//	return err
	//}
	//qb.Select("wait_del_user.video_id").From("wait_del_user").Limit(3)
	//sql := qb.String()
	//_,err = dbops.O.Raw(sql).QueryRows(&videos)
	//if err != nil {
	//	fmt.Println("[ERROR] videoDeleteProducer.QueryRows is err",err)
	//	return err
	//}
	videos,err := dbops.GetWaitDelUsers()
	if err != nil {
		return err
	}

	//将id发送给channel
	for _,v := range videos{
		da <- v.VideoID
	}
	return nil
}

func VideoDeleteConsumers(da DataChan) error {
	//从channel中拿到video_id
	var errMap sync.Map
	var err error
	loop:
	for {
		select {
		case  d := <-da:
			go func(d string) {
				//删除del_video 后再删除本地video文件
				path := strings.Join([]string{"./video"},d)
				err := os.Remove(path)
				if err != nil {
					errMap.Store(d, err)
				}

				//删除数据
				//var newWait dbops.WaitDelUser
				//newWait.VideoID = d
				//
				//_,err = dbops.O.Delete(&newWait)
				err = dbops.DeleteWaitDelUser(d)
				if err != nil {
					errMap.Store(d,err)
				}
			}(d.(string))
		default:
			break loop
		}
	}

	errMap.Range(func(key, value interface{}) bool {
		//只要有一个错误那么就直接返回出错
		err = value.(error)

		if 	value != nil {
			return false
		}
		return true
	})
	return err
}


