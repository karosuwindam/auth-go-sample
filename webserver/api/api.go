package api

import (
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"
	"suth-go-sample/webserver/api/register"
	"suth-go-sample/webserver/api/root"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//
	root.Init()
	login.Init()
	register.Init()

	//
	common.Init("/api/v1", r)
}
