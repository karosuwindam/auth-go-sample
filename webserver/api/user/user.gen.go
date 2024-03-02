package user

import "github.com/gin-gonic/gin"

type ServerIngerface interface {
	PutUserCreate(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserRead(c *gin.Context)
	PostUpdateUserById(c *gin.Context)
	PostUserUpdate(c *gin.Context)
	GetUserReadList(c *gin.Context)
}

func HanlerFormMux(si ServerIngerface, r *gin.RouterGroup) {
	r.PUT("/user", si.PutUserCreate)
	r.POST("/user", si.PostUserUpdate)
	r.DELETE("/user/:id", si.DeleteUser)
	r.GET("/user/:id", si.GetUserRead)
	r.POST("/user/:id", si.PostUpdateUserById)
	r.GET("/user/list", si.GetUserReadList)
}
