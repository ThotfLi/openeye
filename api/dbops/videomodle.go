package dbops

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

func AddVideo(userid uint,name string) error {
	uuid,err := uuid.NewV4()
	if err != nil {
		fmt.Println("[ERROR] uuid invit fail")
		return err
	}

	ctime := 	time.Now().Format("2006-01-02 15:04:05")
	newVideo := VideoInfo{
		VideoInfoId:uuid.String(),
		Name:name,
		User:&User{Id:userid},
		DisplayTime:ctime,
	}

	if p,err := O.Insert(&newVideo); err != nil {
		fmt.Println("[ERROR] Insert fail,p:",p)
		return err
	}

	return nil
}

func GetVideo(videoId string) (*VideoInfo,error){
	newVido := VideoInfo{VideoInfoId:videoId}

	if err := O.Read(&newVido); err != nil {
		fmt.Println("[ERROR] GetVideo fail")
		return nil,err
	}

	return &newVido,nil

}

func DelVideo(videoId string) error {
	newVido := VideoInfo{VideoInfoId:videoId}
	if p,err := O.Delete(&newVido) ; err != nil {
		fmt.Println("[ERROR] DelVideo fail,num p:",p)
		return err
	}
	return nil
}