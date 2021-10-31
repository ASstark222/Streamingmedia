package dbops

import (
	"Streamingmedia/api/defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

/*
对数据库中的session进行操作
TABLE：sessions
session_id TINYTEXT,PRIMARY KEY,NOT NULL
TTL TINYTEXT                              过期时间
login_name VARCHAR(64)
 */

func InserSession(sessionid,username string,ttl int64) error {
	ttlstr := strconv.FormatInt(ttl,10)
	stmtIns,err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,login_name) VALUES (?,?,?)")
	if err != nil{
		return err
	}

	_,err = stmtIns.Exec(sessionid,ttlstr,username)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sessionid string) (*defs.SimpleSeesion,error)  {
	ss := &defs.SimpleSeesion{}
	stmtout,err := dbConn.Prepare("SELECT TTL,login_name FROM sessions WHERE session_id = ?" )
	if err != nil{
		return nil,err
	}

	var ttl string
	var username string
	err = stmtout.QueryRow(sessionid).Scan(&ttl,&username)
	if err != nil&&err != sql.ErrNoRows{
		return nil,err
	}

	if res,err := strconv.ParseInt(ttl,10,64);err == nil{
		ss.TTL = res
		ss.Username = username
	}else{
		return nil,err
	}
	defer stmtout.Close()

	return ss,nil
}

func RetrieveAllSession() (*sync.Map,error)  {
	m := &sync.Map{}
	stmtout,err := dbConn.Prepare("SELECT  * FROM sessions")
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}

	rows,err := stmtout.Query()
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}

	for rows.Next(){
		var id string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id,&ttlstr,&login_name);er != nil{
			log.Printf("retrieve sessions err: %s",err)
			break
		}

		if ttl,err1:=strconv.ParseInt(ttlstr,10,64);err1==nil{
			ss := &defs.SimpleSeesion{TTL: ttl,Username: login_name}
			m.Store(id,ss)
			log.Printf("session id :%s,ttl:%d",id,ss.TTL)
		}
	}
	return m,nil
}

func DeleteSession(sessionid string) error {
	stmtout,err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}

	if _,err := stmtout.Query(sessionid);err !=nil{
		return err
	}

	return nil
}
