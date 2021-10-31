package main

import (
	"Streamingmedia/api/defs"
	"Streamingmedia/api/session"
	"net/http"
)

//验证session和用户

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME= "X-User-Name"

//session验证
func validateUserSession(r *http.Request) bool {
	//获取session id
	sessionid := r.Header.Get(HEADER_FIELD_SESSION)
	//判断session id 是否为空
	if len(sessionid) == 0{
		return false
	}

	//判断session 是否到期
	username,ok := session.IsSessionExpired(sessionid)
	if ok{
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME,username)
	return true
}


//user验证
func ValidateUser(w http.ResponseWriter,r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}