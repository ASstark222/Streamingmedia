package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/",homeHandler)
	router.POST("/",homeHandler)
	router.GET("/userhome",userHomeHandler)
	router.POST("/userhome",userHomeHandler)
	router.POST("/api",apiHandler)

	router.POST("/upload/:vid-id",proxyVideoHandler)

	router.POST("/upload/:vid_id",proxyHandler)

	router.ServeFiles("/statics/*filepath",http.Dir("./template")) //ip:port/statics/template

	return router

}

func main()  {
	r := RegisterHandler()
	http.ListenAndServe(":8080",r)
}