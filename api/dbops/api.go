package dbops

import (
	"Streamingmedia/api/defs"
	"Streamingmedia/api/utils"
	"database/sql"
	"log"
	"time"
)

/*
对数据库中的用户，视频，评论进行操作
进行数据库的增删改查
 */



/*
数据库设计：users

TABLE：users
id USGIGNED INT,PRIMARY KEY,AUTO_INCREMENT
login_name VARCHAR(64),UNIQUE KEY			是用户的唯一标识
pwd TEXT
 */
func AddUserCredential(loginName,pwd string) error {
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?, ?)")
	if err != nil{
		return err
	}
	_,err =stmtIns.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error) {
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users WHENEVER login_Name = ?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil&&err != sql.ErrNoRows{
		return "",err
	}
	defer stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName,pwd string) error {
	stmtDel,err := dbConn.Prepare("DELETE FROM users WHENEVER login_name = ? AND pwd = ?")
	if err != nil{
		log.Printf("DeleteUser error:%s",err)
		return err
	}

	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}



/*video，视频，可以增查删
TABLE:video_info
id VARCHAR(64),PRIMARY KEY,NOT NULL
author_id UNSIGNED INT                 与用户表中的id是同一个
name TEXT
display_ctime TEXT                     页面显示的时间
create_time DATETIME				   视频进入数据库的时间
 */
func AddNewVideo(aid int,name string) (*defs.VideoInfo,error) {
	//uuid
	vid,err := utils.NewUUID()
	if err != nil{
		return nil,err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") //以  月 日 年  时：分：秒  的格式显示时间
	stmtIns,err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime) VALUES (?,?,?,?)")
	if err != nil{
		return nil,err
	}
	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil{
		return nil,err
	}
	res := &defs.VideoInfo{Id: vid,AuthorId: aid,Name: name,DisplayCtime: ctime}
	defer stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo,error) {
	stmtout,err := dbConn.Prepare("SELECT author_id, name , display_ctime FROM video_info WHERE id=?")

	var(
		aid int
		dct string
		name string
	)

	err = stmtout.QueryRow(vid).Scan(&aid,&name,&dct)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}

	if err == sql.ErrNoRows{
		return nil,err
	}

	defer stmtout.Close()

	res := &defs.VideoInfo{Id: vid,AuthorId: aid,Name: name,DisplayCtime: dct}
	return res,nil
}

func DeleteVideInfo(vid string) error {
	stmtDel,err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil{
		return err
	}

	_,err = stmtDel.Exec(vid)
	if err != nil{
		return err
	}

	defer stmtDel.Close()
	return nil
}


/*comments，评论，可以增查
TABLE:comments
id VARCHAR(64),PRIMARY KEY,NOT NULL
video_id VARCHAR(64)
author_id UNSIGNED INT
content TEXT
time DATETIME
*/
func AddNewComments(vid string,aid int,content string) error {
	id,err := utils.NewUUID()
	if err != nil{
		return err
	}
	stmtIns,err := dbConn.Prepare("INSERT INTO comments (id,vide0_id,author_id,content) VALUES (?,?,?,?)")
	if err != nil{
		return err
	}

	_,err= stmtIns.Exec(id,vid,aid,content)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string,from,to int) ([]*defs.Comments,error) {
	stmtout,err := dbConn.Prepare(`SELECT comments.id,users.Login_name, comments.content FROM comments
INNER JOIN users ON comments.author_id = users.id
WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time<=FROM_UNIXTIME(?)`)

	var res []*defs.Comments

	rows,err := stmtout.Query(vid,from,to)
	if err != nil{
		return res,err
	}

	for rows.Next(){
		var id,name,content string
		if err := rows.Scan(&id,&name,&content);err != nil{
			return res,err
		}
		c := &defs.Comments{Id: id,VideoId: vid,Author: name,Content: content}
		res = append(res,c)
	}
	defer stmtout.Close()
	return res,nil
}