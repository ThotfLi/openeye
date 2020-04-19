package dbops

import "time"

type User struct{
	Id uint
	LoginName string `orm:"unique;size(64)"`
	Pwd       string `orm:"size(64)"`
	Videos    []*VideoInfo `orm:"reverse(many)"`
	Commonts  []*Comment   `orm:"reverse(many)"`
}

type VideoInfo struct{
	VideoInfoId string    `orm:"pk;size(64)"`
	User        *User     `orm:"rel(fk)"`
	Name        string    `orm:"type(text)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	DisplayTime string
}

type Comment struct{
	CommentId string `orm:"pk;size(64)"`
	Video     *VideoInfo `orm:"rel(fk)"`
	Author    *User     `orm:"rel(fk)"`
	Content   string    `orm:"type(text)"`
	CreateTime   time.Time `orm:"type(datetime);auto_now_add"`
}

type Sessions struct{
	SessionId string `orm:"size(64);pk"`
	TTL       time.Time `orm:"type(datatime);column(TTL);Null"`  //过期时间
	DataJson  string  `orm:"type(text);NULL"`
}
