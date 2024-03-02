package user

import (
	"strconv"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
)

func (a *API) DeleteUser(c *gin.Context) {
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
		id := 0
		if id, err = strconv.Atoi(c.Param("id")); err != nil {
			c.JSON(400, gin.H{
				"message": "bad request",
			})
			return
		}
		if _, err := users.GetId(id); err == nil && jwtdata.Id != id {
			// ユーザーを削除
			err = users.Delete(id)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "ok",
				})
			} else {
				c.JSON(500, gin.H{
					"message": "internal server error",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"message": "bad request",
			})
		}

	} else {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
	}

}
