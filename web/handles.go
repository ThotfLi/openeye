package web

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"net/http"
	"openeye/session"
)

func homeHandle(w http.ResponseWriter,r *http.Request, params httprouter.Params ) {
	//登录请求
	if r.Method == http.MethodPost {

	}
	sessionid,err := r.Cookie("X-SESSION-ID")
	//没登录并且get请求，返回登录页面
	if err != nil && r.Method == http.MethodGet {
		//尚未登录
		t := template.Must(template.ParseFiles("home.html"))
		t.Execute(w,nil)
	}

	//判断session

	ses := session.Session{SessionID:sessionid.String()}
	err = ses.ReadDB()
	if err != nil {
		//session已过期 需要重新登录
		t := template.Must(template.ParseFiles("home.html"))
		t.Execute(w,nil)
	}

	//已登录
	//返回template
	t := template.Must(template.ParseFiles("userhome.html"))
	t.Execute(w,ses.UserName)

}

func userHomeHandle(w http.ResponseWriter,r *http.Request , params httprouter.Params) {
	ses ,err := r.Cookie("X-SESSION-ID")
	if err != nil {
		http.Redirect(w,r,"/home",302)
	}

	newSes := session.Session{
		SessionID: ses.String(),
	}
	err = newSes.ReadDB()
	if err != nil {
		http.Redirect(w,r,"/login",302)
	}

	t := template.Must(template.ParseFiles("../templates/home.html"))
	t.Execute(w,newSes.UserName)


}

func apiHandle(w http.ResponseWriter,r *http.Request , params httprouter.Params) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"服务器出错",501)
	}
	defer r.Body.Close()

	var newMsg ApiBody
	if err := json.Unmarshal(b,&newMsg); err != nil {
		http.Error(w,"服务器出错",501)
	}


}