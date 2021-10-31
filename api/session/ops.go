package session

import (
	"Streamingmedia/api/dbops"
	"Streamingmedia/api/defs"
	"Streamingmedia/api/utils"
	"sync"
	"time"
)


var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

//从cashe里删除
func deleteExpiredSession(sessionid string)  {
	sessionMap.Delete(sessionid)
	dbops.DeleteSession(sessionid)
}

//从数据库中加载
func LoadSessionFromDB()  {
	r,err := dbops.RetrieveAllSession()
	if err !=nil{
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSeesion)
		sessionMap.Store(key,ss)
		return true
	})
}
//生成新的
func GenerateNewSessionId(username string ) string {
	id,_:= utils.NewUUID()
	ct :=time.Now().UnixNano()/100000  //精确到毫秒
	ttl := ct + 30*60*1000  //过期时间：30min

	ss := &defs.SimpleSeesion{TTL: ttl,Username: username}
	sessionMap.Store(id,ss)//  session的uuid和内容对应
	dbops.InserSession(id,username,ttl)
	return id
}
//判断是否过期
func IsSessionExpired(sessionid string) (string,bool) {
	ss,ok := sessionMap.Load(sessionid)
	if ok {
		ct := time.Now().UnixNano()/100000
		if ss.(*defs.SimpleSeesion).TTL<ct{
			//删除
			deleteExpiredSession(sessionid)
			return "",true
		}
		return ss.(*defs.SimpleSeesion).Username,false
	}
	return "",true
}