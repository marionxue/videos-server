package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marionxue/videos_server/api/defs"
	"github.com/marionxue/videos_server/api/utils"
	"log"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO t_users(login_name,pwd) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd from t_users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err

	}
	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from t_users where login_name=? and pwd =?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	ctime := time.Now().Format("Jan 02 2006,15:04:05")
	stmtIns, err := dbConn.Prepare(`insert into t_video_info
		(id,author_id,name,display_ctime) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id:vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil

}


func GetVideoInfo(vid int)(*defs.VideoInfo,error){
	stmtOut,err:= dbConn.Prepare(`select author_id,name,display_name,create_time from t_videos_info where id=?`)

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid,&name,&dct)
	if err!=nil&& err!=sql.ErrNoRows{
		return nil,err
	}

	if err==sql.ErrNoRows{
		return nil,nil
	}
	defer stmtOut.Close()

	res:= &defs.VideoInfo{Id:vid,AuthorId:aid,DisplayCtime:dct,Name:name}
	return res,nil
}

func DeleteVideoInfo(vid int) error{
	stmtDel,err:= dbConn.Prepare("delete from t_videos_info where id=?")
	if err!=nil{
		return nil
	}

	_,err = stmtDel.Exec(vid)
	if err !=nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}