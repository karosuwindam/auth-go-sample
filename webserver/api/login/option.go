package login

import "github.com/gin-gonic/gin"

// ginによる許可確認
func (a *API) OptionLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "X-POST, Content-Type, Authorization")
	c.Header("Access-Control-Max-Age", "86400")
	c.JSON(204, gin.H{
		"message": "ok",
	})
}
