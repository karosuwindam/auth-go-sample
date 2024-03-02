package api

import (
	"suth-go-sample/webserver/api/hello"
	"suth-go-sample/webserver/api/login"
	"suth-go-sample/webserver/api/user"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	hello.HanlerFormMux(hello.NewAPI(), v1)
	login.HanlerFormMux(login.NewAPI(), v1)
	user.HanlerFormMux(user.NewAPI(), v1)
}
