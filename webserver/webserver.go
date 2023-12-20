package webserver

import (
	"net/http"
	"suth-go-sample/config"
	"suth-go-sample/webserver/api"
	"suth-go-sample/webserver/webconfig"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// ginによるwebサーバの初期化
func Init() error {
	r = gin.Default()
	// r.Use(setHeader()) // 共通ヘッダーの設定
	setCORS(r)
	api.Init(r)                              // apiの初期化
	if err := webconfig.Init(); err != nil { //Webサーバ周辺の初期化
		return err
	}
	return nil
}

func Start() {
	r.Run(config.Web.Host + ":" + config.Web.Port)
}

// ginのCORSの設定
func setCORS(r *gin.Engine) {
	allowOrigins := config.Web.AllowOrigin
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowOrigins},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
