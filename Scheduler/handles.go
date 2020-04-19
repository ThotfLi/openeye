package Scheduler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"openeye/Scheduler/dbops"
)

func VidDelHandle (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoID := p.ByName("videoID")

	err := dbops.InsertWaitDelUser(videoID)
	if err != nil {
		http.Error(w,"服务器错误",501)
	}
}