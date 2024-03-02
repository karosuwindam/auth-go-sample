package user

import (
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
)

// ユーザーを作成する
func (a *API) PutUserCreate(c *gin.Context) {
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
	// 管理者以上の権限を持っているか確認
	if userData.Authority >= common.ADMIN {
		var user CreateUser
		c.BindJSON(&user)
		if user.Name == "" || user.Password == "" {
			c.JSON(400, gin.H{
				"message": "bad request",
			})
			return
		}
		if err := hashPassword(&user); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
		if err := users.Add(user.Name, user.Password, user.Authority); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "registered",
		})
	} else {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
	}
}
