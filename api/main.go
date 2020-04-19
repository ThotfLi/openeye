package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"openeye/session"
)
type middleWareHandler struct {
	r *httprouter.Router
}

type Request struct{
	Session session.ISession
	r       *http.Request
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {

	m.r.ServeHTTP(w,r)
}


func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.GET("/user/:username",Login)
	return router
}

func main() {
	r := RegisterHandles()
	http.ListenAndServe(":8000", r)
}
