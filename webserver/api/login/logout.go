package login

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ginによるログアウト処理
func (a *API) PostLogout(c *gin.Context) {
	//ヘッダからトークンを取得
	tokenString := c.Request.Header.Get("Authorization")
	output := UserMessage{}
	//jwt情報を確認
	if tokenString != "" {
		if j, err := UnpackJwt(tokenString); err != nil {
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
