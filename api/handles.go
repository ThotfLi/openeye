package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"openeye/session"
	"openeye/api/dbops"
	"time"
)

var (
	GET_SESSION_ID = "X-SESSION-ID"
)
func CreateUser(w http.ResponseWriter,r *http.Request,param httprouter.Params){
}

func Login(w http.ResponseWriter,r *http.Request,param httprouter.Params) {
	//客户端"是否存在"session，存在则判断"是否过期"没过期直接返回。过期、不存在继续向下执行
	if c,err := r.Cookie(session.GET_SESSION_COOKIE);err == nil {
		if b,_ := session.SesMangerObj.IsExpiration(c.Value); !b {
			//session未过期直接返回，不继续登录任务
			io.WriteString(w,"请不要重复登录")
			return
		}
	}

	//验证用户名、密码
	//if username == "skyleiou" {
	//
	//	//验证通过
	//	//新建session
	//	newsession := session.NewSession(username)
	//	newsession.SetAttr("username", username)
	//	newsession.SetOverdueTime(10)
	//	//将新的session写入数据库
	//	newsession.WriteDB()
	//	//将新的session添加到session管理器
	//	session.SesMangerObj.AddMap(newsession.GetSessionID(), newsession)
	//
	//	//将session写入Cookie
	//	c1 := http.Cookie{Name: session.GET_SESSION_COOKIE,
	//		Value: newsession.GetSessionID()}
	//	http.SetCookie(w, &c1)
	//	//跳转登录页面
	//
	//	//验证未通过继续登录
	//}

	username := r.FormValue("username")
	pwd      := r.FormValue("pwd")

	ppwd ,err := dbops.GetUser(username)
	if err != nil {
		http.Error(w,"用户或密码错误",400)
	}

	if pwd == ppwd {
		ses := session.NewSession(username)
		c := http.Cookie{
			Name:      GET_SESSION_ID,
			Value:      ses.GetSessionID(),
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(w,&c)
		ses.WriteDB()
		return
	}


}

//如果尚未登录、登录超时 则禁止访问handle
func IsLoginIN(handle httprouter.Handle) httprouter.Handle{
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		c,err := r.Cookie(session.GET_SESSION_COOKIE)
		b,_ := session.SesMangerObj.IsExpiration(c.Value)
		if err != nil || b {
			//cookie不存在，尚未登录
			//设置跳转登录页面
			w.Header().Set("Location","http://127.0.0.1/login")
			w.WriteHeader(302)
			return
		}

		handle(w,r,params)
	}
}