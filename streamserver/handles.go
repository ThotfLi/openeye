package streamserver

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetVideo(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	//vid := param.ByName("videoID")
	//a,err := os.Open(`F:\迅雷下载\小丑回魂.mp4`)
	a,err := os.Open(`E:\迅雷下载\hhbb_zcm@www.sis001.com@081618_729-1pon\081618_729-1pon.mp4`)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),a)

}

func UpVideo(w http.ResponseWriter,r *http.Request,param httprouter.Params){
	r.Body = http.MaxBytesReader(w,r.Body,UPFILE_MAXSIZE)
	if err :=r.ParseMultipartForm(UPFILE_MAXSIZE);err != nil{
			http.Error(w,"file is big",404)
		}
	file,_,err := r.FormFile("file")
	if err != nil {
		http.Error(w,"server 异常",501)
	}
	data ,err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w,"server 异常",501)
	}

	fn := param.ByName("videoID")
	err = ioutil.WriteFile(VIDEO_DIR+fn,data,066)
	if err != nil {
		http.Error(w,"SERVER 异常",501)
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"uploade successfull")

}


