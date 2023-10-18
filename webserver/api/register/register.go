package register

import (
	"suth-go-sample/config"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
}

// ginによるユーザの登録
func RegisterPut(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	// ここでuserをDBに登録する処理を書く
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

// パスワードの更新
func UpdatePassword(c *gin.Context) {
	//ヘッダからトークンを取得
	jwtdata, err := login.GetJwtUser(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}
	//ユーザー情報からパスワードを更新する
	var user User
	c.BindJSON(&user)
	if user.Name == "" || user.Password == "" {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	u, err := users.Get(user.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if u.Name == jwtdata.Name || u.Authority >= common.ADMIN {
		user.Authority = u.Authority
		if err := hashPassword(&user); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
		if err := users.Update(int(u.Model.ID), user.Password); err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
			})
			return
		}
	} else {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "updated",
	})
}

// パスワードをペッパー文字列でハッシュ化する
func hashPassword(user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Name+user.Password+config.JWT.Pepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}
