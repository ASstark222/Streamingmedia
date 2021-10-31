package main

import (
	"io"
	"net/http"
)

func sendResopnse(w http.ResponseWriter,sc int,resp string)  {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}
