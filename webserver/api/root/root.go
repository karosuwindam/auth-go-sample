package root

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Origin", string(c.Request.Host))
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
