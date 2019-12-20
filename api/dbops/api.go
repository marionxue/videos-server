package dbops

import "database/sql"

func Openconn() *sql.DB {}

func AddUserCredential(loginName string, pwd string) error {}

func GetUserCredential(loginName string) (string, error) {}