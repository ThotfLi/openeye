package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"streamserver/abc"
	"io"
)

type  streamserver struct {
	stream  *httprouter.Router
	l        abc.ConnLimiter
}
func NewStreamServer(conNum uint) streamserver {
	return streamserver{l:abc.NewConnLimiter(conNum)}
}

func (s streamserver)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if b := s.l.AddConn() ; !b {
		io.WriteString(w,"满了")
		return
	}
	s.stream.ServeHTTP(w,r)
	s.l.ReduceConn()
}

func RegisterHandle()*httprouter.Router{
	router := httprouter.New()
	router.GET("/video/:videoID",abc.GetVideo)
	router.POST("/video/:videoID",abc.UpVideo)
	return router
}

func main(){
	r := RegisterHandle()
	s := NewStreamServer(2)
	s.stream = r
	http.ListenAndServe(":8888",s)
}

