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
	r.Use(setHeader()) // 共通ヘッダーの設定
	api.Init(r)        // apiの初期化

}

func Start() {
	r.Run(config.Web.Host + ":" + config.Web.Port)
}

// 共通ヘッダーの設定
func setHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.Web.AllowOrigin)
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	}
}
