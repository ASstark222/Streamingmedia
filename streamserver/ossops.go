package main

import (
	"Streamingmedia/streamserver/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var EP string
var SK string
var AK string

func init()  {
	AK = "LTAI5tE7cDCkcYzyUHobTcpF"
	SK = "WMHiQPY9Z4w0sYQnFQgmnafTKlL1Ob"
	EP = config.GetOssAddr()

}

func UploadToOss(filename string,path string,bn string) bool {
	client,err := oss.New(EP,AK,SK)
	if err != nil{
		log.Println("Init oss error: %s ",err)
		return false
	}

	bucket,err := client.Bucket(bn)
	if err != nil{
		log.Println("bucket err: %s ",err)
		return false
	}

	err = bucket.UploadFile(filename,path,500*1024,oss.Routines(3))
	if err != nil{
		log.Println("uploading object err :%s",err)
		return false
	}
	return true
}