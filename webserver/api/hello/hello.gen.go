package hello

import "github.com/gin-gonic/gin"

type ServerIngerface interface {
	GetHello(c *gin.Context)
}

func HanlerFormMux(si ServerIngerface, r *gin.RouterGroup) {
	r.GET("/hello", si.GetHello)
}
