package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// 注册用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// 登录 Method: POST, SC: 200,400,500
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
