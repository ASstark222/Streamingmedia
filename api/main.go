package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*
handler -->  validationn{1.request,2.user} -->bussiness logic  --> response
 */

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

type middleWareHandler struct {
	r *httprouter.Router
}
func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w,r)
}



func main()  {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000",mh)
}