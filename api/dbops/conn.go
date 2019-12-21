package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init(){
	dbConn,err=sql.Open("mysql","root:admin123@tcp(127.0.0.1:3306)/videos_server?charset=utf8mb4")
	if err!=nil{
		panic(err.Error())
	}
}