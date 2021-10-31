package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*
将video格式化成二进制的比特流
用stream的方式发送到客户端
 */
func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR+vid

	video,err := os.Open(vl)
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"error")
		return
	}

	w.Header().Set("Content-type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)

	defer video.Close()

}

/*

 */
func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);err!=nil{
		sendErrorResponse(w,http.StatusBadRequest,"file is too big")
		return
	}

	file,_,err := r.FormFile("file") //http 的form中 name = "file"
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"Internal server")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err != nil{
		log.Printf("Read file error:%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn,data,0666)
	if err != nil{
		log.Printf("writer file error :%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}


	ossfn := "videos/"+ fn
	path := "./videos/"+fn
	bn := "govideos"
	ret := UploadToOss(ossfn,path,bn)
	if !ret{
		sendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}

	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"upload succeed")
}
