package user

import (
	"suth-go-sample/config"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
}

// ユーザーを作成する
func Create(c *gin.Context) {
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
	}
	c.JSON(401, gin.H{
		"message": "unauthorized",
	})
}

// パスワードをペッパー文字列でハッシュ化する
func hashPassword(user *CreateUser) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Name+user.Password+config.JWT.Pepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

// パスワードが一致するかを確認する
func checkPassword(user *CreateUser) bool {
	if u, err := users.Get(user.Name); err != nil {
		return false
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Name+user.Password+config.JWT.Pepper)); err != nil {
			return false
		}
	}
	return true
}
