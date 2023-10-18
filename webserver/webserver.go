package webserver

import (
	"suth-go-sample/config"
	"suth-go-sample/webserver/api"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// ginによるwebサーバの初期化
func Init() {
	r = gin.Default()
	api.Init(r) // apiの初期化

}

func Start() {
	r.Run(config.Web.Host + ":" + config.Web.Port)
}
