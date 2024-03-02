package login

import "github.com/gin-gonic/gin"

type ServerIngerface interface {
	OptionLogin(c *gin.Context)
	GetLogin(c *gin.Context)
	PostLogin(c *gin.Context)
	PostLogout(c *gin.Context)
}

func HanlerFormMux(si ServerIngerface, r *gin.RouterGroup) {
	r.OPTIONS("/login", si.OptionLogin)
	r.OPTIONS("/logout", si.OptionLogin)
	r.GET("/login", si.GetLogin)
	r.POST("/login", si.PostLogin)
	r.POST("/logout", si.PostLogout)
}
