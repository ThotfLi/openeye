package main

import (
	"io"
	"net/http"
	"openeye/api/defs"
)

func sendErrorResponse(w http.ResponseWriter,err defs.ErrorResponse){
	w.WriteHeader(err.HttpSC)
	io.WriteString(w,err.Error.Error)
}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int){
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}


