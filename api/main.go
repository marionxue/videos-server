package main

import (
	"github.com"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", github_com.CreateUser)
	router.POST("/user/:user_name", github_com.Login)

	return router
}
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}


