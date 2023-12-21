package user

import "github.com/gin-gonic/gin"

func Update(c *gin.Context) {
	// ToDo: ユーザー情報を更新する Admin用
	c.JSON(200, gin.H{
		"message": "update",
	})
}

func UpdateById(c *gin.Context) {
	// ToDo: Idを指定してユーザー情報を更新する UserとAdmin用
	// この機能はパスワードのみ更新可能
	c.JSON(200, gin.H{
		"message": "update",
	})
}
