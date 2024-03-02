package user

import (
	"strconv"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
)

// IDを指定してユーザー情報を取得する
func (a *API) GetUserRead(c *gin.Context) {
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
	output := User{Id: id}
	if tmp, err := users.GetId(id); err == nil {
		output.Name = tmp.Name
		output.Auth = tmp.Authority
	}
	if id == jwtdata.Id || userData.Authority >= common.ADMIN {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    output,
		})
		return
	}
	c.JSON(401, gin.H{
		"message": "unauthorized",
	})
}

func (a *API) GetUserReadList(c *gin.Context) {
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
	if userData.Authority >= common.ADMIN {
		if userList, err := makeUserList(); err == nil {
			c.JSON(200, gin.H{
				"message": "ok",
				"data":    userList,
			})
			return
		}
	}
	c.JSON(401, gin.H{
		"message": "unauthorized",
	})

}
