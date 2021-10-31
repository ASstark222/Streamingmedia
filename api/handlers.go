package main

import (
	"Streamingmedia/api/dbops"
	"Streamingmedia/api/defs"
	"Streamingmedia/api/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

/*
处理器
不同的URL定位交由不同的处理器处理
 */
func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//从request里把body读出来
	res,_:=ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res,ubody);err != nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd);err !=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true,SessionId: id}
	if resp,err := json.Marshal(su);err != nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(w,string(resp),201)
	}
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}

