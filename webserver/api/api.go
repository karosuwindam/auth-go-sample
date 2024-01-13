package api

import (
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"
	"suth-go-sample/webserver/api/logout"
	"suth-go-sample/webserver/api/root"
	"suth-go-sample/webserver/api/user"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	common.Init("/api/v1", r,
		root.Init,
		login.Init,
		// register.Init,
		logout.Init,
		user.Init,
	)
}
