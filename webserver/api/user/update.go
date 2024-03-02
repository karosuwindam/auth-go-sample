package user

import (
	"strconv"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
)

func (a *API) PostUserUpdate(c *gin.Context) {

	//ヘッダからトークンを取得
	jwtdata, err := login.GetJwtUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}
	// トークン情報からIDを指定してテーブルからユーザ情報を取得
	userData, err := users.GetId(jwtdata.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	//受け取ったJSONからユーザー情報を取得
	var updateUser UpdateUser
	c.BindJSON(&updateUser)

	if userData.Authority >= common.ADMIN {
		flag := false
		// ユーザー情報を更新する Admin用
		if flag, err = updateUserData(&updateUser); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
		if flag {
			c.JSON(200, gin.H{
				"message": "update",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "no update",
			})
		}
	} else {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
	}

}

func (a *API) PostUpdateUserById(c *gin.Context) {
	// ToDo: Idを指定してユーザー情報を更新する UserとAdmin用

	//ヘッダからトークンを取得
	jwtdata, err := login.GetJwtUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}
	id := 0
	//idの文字列を数列に変換する
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	// トークン情報からIDを指定してテーブルからユーザ情報を取得
	userData, err := users.GetId(jwtdata.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	//受け取ったJSONからユーザー情報を取得
	var newpass NewPassword
	c.BindJSON(&newpass)
	if newpass.NewPass == "" {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	// この機能はパスワードのみ更新可能
	if id == jwtdata.Id || userData.Authority >= common.ADMIN {
		if err := updatePassword(id, &newpass); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "update",
			})
		}
	} else {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
	}
}
