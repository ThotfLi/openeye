package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ReigistionRouter()*httprouter.Router{
	router := httprouter.New()

	router.GET("/home",homeHandle)
	router.POST("/home",homeHandle)

	router.GET("/userhome",userHomeHandle)
	router.POST("/userhome",userHomeHandle)

	router.POST("/api",apiHandle)
	router.ServeFiles("/static/*filepath",http.Dir("../templates"))
	return router
}
