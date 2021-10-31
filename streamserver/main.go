package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
/*
利用中间件加上流控功能
 */
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetCon(){
		sendErrorResponse(w,http.StatusTooManyRequests,"Too Many Request")
		return
	}

	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}

func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler  {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router
}

func main()  {
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r,2)
	http.ListenAndServe(":9000",mh)
}