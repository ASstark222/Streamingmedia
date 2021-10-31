package defs

/*
数据结构定义
 */

//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}
//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

//Date model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

//comments
type Comments struct {
	Id string
	VideoId string
	Author string
	Content string
}

//session
type SimpleSeesion struct {
	Username string //用户登录名
	TTL int64
}