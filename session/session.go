package session

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
	"openeye/api/dbops"
)

//每个用户具有独立的session对象
//使用sync.map 安全的进行并发设置session属性

//使用的时候如果不设置过期时间，那默认就在关闭浏览器后过期
//建议使用SetOverdueTime方法设置过期时间

type ISession interface{
	//设置过期时间
	SetOverdueTime(int)
	//获取过期时间
	GetOverdueTime()                  string

	//设置属性

	SetAttr(key string,value interface{})  error
	//获取属性
	GetAttr(string)                   (interface{},error)
	//删除属性
	DelAttr(string)

	//获取sessionID
	GetSessionID()                    string

	//返回true则时间已过期
	ISExpiration()                    bool

	//属性数据转换成JSON
	DataToJson()                      (string,error)

	//将当前session中的数据写入、更新到数据库
	WriteDB ()                        error

	//将当前sessionID的数据从数据库读到当前session对象
	ReadDB()                          error

	GetSyncMap ()                     *sync.Map
}

type Session struct{
	SessionID  string
	sycMap     *sync.Map
	TTL        time.Time  //过期时间的时间戳
	UserName   string
}

func NewSession(username string)ISession {
	ssionID,_ := uuid.NewV4()
	strssionID := ssionID.String()
	return &Session{
		SessionID:strssionID,
		sycMap:   &sync.Map{},
		UserName:username,
	}
}
//输入的是秒
func (ses Session) GetUserName() string{
	return ses.UserName
}

func (ses *Session) SetOverdueTime (s int){
	t  := time.Now()
	z  := t.Add(time.Duration(s)*time.Second)
	ses.TTL = z
}

func (ses *Session) GetOverdueTime () string {
	ts := ses.TTL.Format("2006/01/02 15:04:05")
	return ts
}

func (ses *Session) SetAttr (k string,v interface{}) error {
	if _,ok := ses.sycMap.Load(k);ok {
		fmt.Println("[ERROR] Set Attr fail")
		return errors.New("Attr set fail")
	}
	ses.sycMap.Store(k,v)
	return nil
}

func (ses  Session) GetAttr (k string) (interface{},error) {
	if v,ok := ses.sycMap.Load(k);ok{
		return v,nil
	}
	return nil ,errors.New("Get Attr session")
}

func (ses  Session)	GetSessionID() string{
	return ses.SessionID
}

func (ses  *Session) GetSyncMap ()*sync.Map{
	return ses.sycMap
}

func (ses  Session) ISExpiration () bool {
	nowT := time.Now()

	//时间已过期
	if ses.rTTL().Before(nowT){
		return true
	}
	return false
}

func (ses  Session) rTTL () time.Time{
	return ses.TTL
}

func (ses *Session) WriteDB ()error {
	var newSes dbops.Sessions
	newSes.SessionId = ses.GetSessionID()
	err := dbops.O.Read(&newSes)

	//此sessionID不存在，则创建一条新的session数据
	if err != nil {
		newSes.TTL = ses.TTL
		newSes.DataJson,err = ses.DataToJson()
		if err != nil {
			fmt.Println("[ERROR] The Session.WriteDB() error")
		}
		if _,err := dbops.O.Insert(&newSes);err != nil {
			fmt.Println("[ERROR] In the Session.Write.DB(),db.Insert err",err)
			return err
		}
	}

	//此session已存在，则对字段TTL、DataJson 进行更新
	if err != nil {
		if newSes.DataJson,err = ses.DataToJson();err != nil {
			fmt.Println("[ERROR] In the Session.WriteDB(), DataToJson() is err:",err)
			return err
		}
		newSes.TTL = ses.rTTL()
		if _,err := dbops.O.Update(&newSes); err != nil {
			fmt.Println("[ERROR]In the Session.WriteDB(),db.Upate() is err:",err)
			return err
		}

	}
	return nil
}

func (ses *Session) DataToJson() (string,error){
	newMap := make(map[string]interface{})

	f := func(key interface{},value interface{}) bool {
		kkey := key.(string)
		newMap[kkey] = value
		return true
	}
	ses.GetSyncMap().Range(f)

	mapJson,err := json.Marshal(newMap)
	if err != nil {
		return "",err
	}
	return string(mapJson),nil
}

func (ses *Session) ReadDB() error {
	//从数据库拿到session
	newSession := dbops.Sessions{
		SessionId: ses.SessionID,
	}
	err := dbops.O.Read(&newSession)
	if err != nil {
		fmt.Println("[ERROR] Session.ReadDB() is err ")
		return err
	}

	//查看session是否过期
	if ses.ISExpiration(){
		return SESSION_EXPIRATION
	}

	//session没过期 将当前Session的数据、TTL 初始化
	ses.TTL = newSession.TTL
	m := make(map[string]interface{})
	json.Unmarshal([]byte(newSession.DataJson),m)
	//循环m将属性写入session的sync.map
	for k,v := range m{
		ses.GetSyncMap().Store(k,v)
	}
	return nil

}

func (ses *Session) DelAttr(k string){
	ses.GetSyncMap().Delete(k)
}