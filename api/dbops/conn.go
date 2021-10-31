package dbops


//用于数据库链接
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"  //加载mysql的驱动
)

var (
	//全局变量.
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn,err =sql.Open("mysql","root:1029384756qpwo@#@tcp(127.0.01:3306)/video_server?charset=utf8")
	if err != nil{
		panic(err.Error())
	}
}
