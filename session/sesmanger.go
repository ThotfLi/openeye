package session

import (
	"fmt"
	"net/http"
)

//session管理器
//保存和管理全局session
//管理和数据库的通信

//加载本模块即可init SesManger的实例SesMangerObj
//操作SesMangerObj即可
type ISesManger interface{
	//获取一个map
	GetSession(sessionID string) (ISession,error)
	//添加一个map
	AddMap(sessionID string, ses ISession) error
	//session是否过期
	//bool = true 已过期
	IsExpiration(sessionID string) (bool,error)
}

func init(){
	SesMangerObj = NewSesManger()
}
var SesMangerObj ISesManger

type SesManger struct{
	m   map[string]ISession
}
var(
	GET_SESSION_COOKIE = "X-SESSION-ID"
)
func NewSesManger () ISesManger {
	return &SesManger{m:make(map[string]ISession)}
}

func (ses *SesManger)GetSession(sessionID string)(ISession,error){
	if synmap,ok :=  ses.returnM()[sessionID];ok{
		return synmap,nil
	}
	fmt.Println("[ERROR]sessionid not find，",sessionID)
	return nil,SESSION_ISNOTFIND

}

func (ses *SesManger)returnM()map[string]ISession{
	return ses.m
}

func (ses *SesManger)AddMap(sessionID string,newsesion ISession)error{
	  //判断session是否存在
	  if _,ok := ses.returnM()[sessionID]; ok {
	  	return SESSION_ISEXISTENCE
	  }
	  //不存在则添加session
	  m := ses.returnM()
	  m[sessionID] = newsesion
	  return nil
}

func (ses *SesManger)IsExpiration(sessionID string) (bool,error){
	if s,err := ses.GetSession(sessionID);err != nil {
		//没找到session也按照过期处理
		return true,SESSION_ISNOTFIND
	}else {
		return s.ISExpiration(),nil
	}
}

//从cookie拿到sessionID 并返回相应的session对象
func (ses *SesManger)SessionStart(w http.ResponseWriter,r *http.Request) (ISession,error) {
	//如果session 存在那直接返回一个session对象
	c ,err := r.Cookie(GET_SESSION_COOKIE)
	if err != nil {
		return nil ,err
	}

	//获取session
	isession,err := ses.GetSession(c.Value)
	if err != nil {
		return nil ,err
	}

	//判断session是否过期
	if isession.ISExpiration(){
		return nil ,SESSION_EXPIRATION
	}


	return isession,nil
}