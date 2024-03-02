package hello

import "github.com/gin-gonic/gin"

func (a *API) GetHello(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
