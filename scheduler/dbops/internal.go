package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
TABLE:video_del_rec   //视频删除记录
video_id varchar(64) primary key

api get videoid from mysql
dispatcher get videoid from mysql  send to datachannel
executor get videoid from datachannel --> delete videos
 */


//读出需要删除的video id
func ReadVideoDeletionRecord(count int) ([]string,error) {
	stmtout,err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")

	var ids []string

	if err != nil{
		return ids,err
	}

	rows,err :=stmtout.Query(count)
	if err != nil{
		log.Printf("Query VideoDeletionRecord err;%s",err)
		return ids,err
	}

	for rows.Next(){
		var id string
		if err := rows.Scan(&id);err !=nil{
			return ids,err
		}

		ids = append(ids,id)
	}

	defer stmtout.Close()
	return ids,nil
}


//删除Vidoe
func DelVideoDeletionRecord(vid string) error {
	stmtDel,err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id = ?")
	if err !=nil{
		return err
	}

	_,err = stmtDel.Exec(vid)
	if err !=nil{
		log.Printf("Deleting video err : %s",err)
		return err
	}

	defer stmtDel.Close()
	return nil
}