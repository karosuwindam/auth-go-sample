package logout

import (
	"fmt"
	"log"
	"suth-go-sample/webserver/api/login"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	ext  int64  `json:"ext"`
}

// ginによる許可確認
func LoginOption(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "X-POST, Content-Type, Authorization")
	c.Header("Access-Control-Max-Age", "86400")
	c.JSON(204, gin.H{
		"message": "ok",
	})
}

// ginによるログアウト処理
func LogoutPost(c *gin.Context) {
	//ヘッダからトークンを取得
	tokenString := c.Request.Header.Get("Authorization")
	output := login.UserMessage{}
	//jwt情報を確認
	if tokenString != "" {
		if j, err := login.UnpackJwt(tokenString); err != nil {
			log.Println(err.Error())
		} else {
			//logout処理
			fmt.Printf("logout name=%v,id=%v\n", j.Name, j.Id)
		}
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    output,
	})
}
