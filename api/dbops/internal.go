package dbops

import (
	"database/sql"
	"log"
	"sync"
	"videos_server/api/defs"
)

func InserSession(sid string, ttl int64, uname string) error {
	//ttlstr:=strconv.FormatInt(ttl,10)
	stmtIns, err := dbConn.Prepare(`insert into t_sessions(id,ttl,login_name) values (?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttl, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetriveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}

	stmtOut, err := dbConn.Prepare(`select ttl,login_name from t_sessions where id=?`)
	if err != nil {
		return nil,err
	}
	var ttl int64
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err != nil && err!= sql.ErrNoRows{
		return nil,err
	}

	ss.TTL = ttl
	ss.Username=uname
	defer stmtOut.Close()
	return ss,nil


}

func RetriveAllSessions()(*sync.Map,error){
	m:=&sync.Map{}
	stmtOut,err:=dbConn.Prepare("select * from t_sessions")
	if err!=nil{
		log.Printf("%s",err)
	}
	rows,err:= stmtOut.Query()
	if err!=nil{
		log.Printf("%s",err)
		return nil,err
	}

	for rows.Next(){
		var id string
		var ttl int64
		var login_name string
		if err:= rows.Scan(&id,&ttl,&login_name);err!=nil{
			log.Printf("retrive sessions error: %s",err)
			break
		}

		ss:=&defs.SimpleSession{Username:login_name,TTL:ttl}
		m.Store(id,ss)
		log.Printf("session id: %s,ttl: %d",id,ss.TTL)
	}
	return m,nil
}

func DeleteSession(sid string) error{
	stmtOut,err:=dbConn.Prepare("DELETE FROM t_sessions where id=?")
	if err!=nil{
		log.Printf("%s",err)
		return err
	}
	if _,err:= stmtOut.Query(sid);err!=nil{
		return err
	}
	return nil
}