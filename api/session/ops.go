package session

import (
	"sync"
	"time"
	"videos_server/api/dbops"
	"videos_server/api/defs"
	"videos_server/api/utils"
)

func AddSession() {

}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64{
	return time.Now().UnixNano()/1000000
}

func LoadSessionsFromDB() {
	r, err := dbops.RetriveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id,_:= utils.NewUUID()
	ct:=nowInMilli()
	ttl:=ct+30*60*1000//serverside seesion valid time:30min
	ss:=&defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InserSession(id,ttl,un)

	return id

}

func deleteExpireSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// nil,true
func IsSessionExpire(sid string) (string, bool) {
	ss,ok:= sessionMap.Load(sid)
	if ok{
		ct:= nowInMilli()
		if ss.(*defs.SimpleSession).TTL< ct{
			deleteExpireSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}
