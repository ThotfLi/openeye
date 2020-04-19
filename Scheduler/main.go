package Scheduler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewhandelServe()*httprouter.Router{
	r := httprouter.New()
	r.GET("/video-delete-record/:videoID",VidDelHandle)
	return r
}

func main(){
	r := NewhandelServe()
	http.ListenAndServe("9090",r)
}