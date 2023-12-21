package user

import "github.com/gin-gonic/gin"

func Delete(c *gin.Context) {
	//ToDo: ユーザー情報を削除する Admin用
	c.JSON(200, gin.H{
		"message": "delete",
	})
}
