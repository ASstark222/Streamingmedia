package main

import (
	"Streamingmedia/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

/*
发送返回消息
 */
//发送错误响应
func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrResponse)  {
	w.WriteHeader(errResp.HttpSC)
	resStr,_ := json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))
}
//发送正常响应
func sendNormalResponse(w http.ResponseWriter,resp string,sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}