package login

import (
	"errors"
	"fmt"
	"suth-go-sample/config"
	"suth-go-sample/tables/users"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// ginによるログイン確認
func LoginGet(c *gin.Context) {
	//ヘッダからトークンを取得
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"token":   "",
		})
		return
	}
	//jwt再発行する
	tokenString, err := reissueJwt(tokenString)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"token":   "",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"token":   tokenString,
	})
}

// ginによるログイン処理
func LoginPost(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	if !CheckPassword(&user) {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"token":   "",
		})
		return
	}
	token, err := GenerateJWT(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"token":   "",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"token":   token,
	})
}

// ユーザー名とパスワードが一致するかを確認する
func CheckPassword(user *User) bool {
	if user.Name == "" || user.Password == "" {
		return false
	}
	tmp := User{}
	if u, err := users.Get(user.Name); err != nil {
		return false
	} else {
		tmp.Name = u.Name
		tmp.Password = u.Password
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tmp.Password), []byte(user.Name+user.Password+config.JWT.Pepper)); err != nil {
		return false
	}
	return true
}

// user情報からjwtを生成する
func GenerateJWT(user *User) (string, error) {
	u, err := users.Get(user.Name)
	if err != nil {
		return "", err
	}
	//jwtを生成する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.Model.ID,
		"name": u.Name,
		"ext":  time.Now().Add(time.Hour * time.Duration(config.JWT.LifeTimeHour)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWTコードのデコード
func UnpackJwt(jwtdata string) (JwtData, error) {
	//Bearerの存在を確認してあったら削除
	if len(jwtdata) > 7 {
		if jwtdata[:7] == "Bearer " {
			jwtdata = jwtdata[7:]
		}
	}
	var output JwtData
	token, err := jwt.Parse(jwtdata, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT.Secret), nil
	})
	if err != nil {
		return output, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		output.Id = int(claims["id"].(float64))
		output.Name = claims["name"].(string)
		output.ext = int64(claims["ext"].(float64))
	} else {
		return output, errors.New("token is invalid")
	}
	return output, nil
}

// JWTの有効期限を確認
func CheckJwt(jwtdata string) bool {
	//Bearerを削除
	jwtdata = jwtdata[7:]
	//jwtをデコード
	tmp, err := UnpackJwt(jwtdata)
	if err != nil {
		return false
	}
	//有効期限を確認
	if tmp.ext < time.Now().Unix() {
		return false
	}
	return true
}

// JWTの再発行
func reissueJwt(jwtdata string) (string, error) {
	//jwtの有効期限を確認
	if !CheckJwt(jwtdata) {
		return "", errors.New("token is invalid")
	}
	//jwtを生成する
	tmp, _ := UnpackJwt(jwtdata)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   tmp.Id,
		"name": tmp.Name,
		"ext":  time.Now().Add(time.Hour * time.Duration(config.JWT.LifeTimeHour)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWTからユーザー情報を取得
func GetJwtUser(c *gin.Context) (JwtData, error) {
	jwtdata := c.Request.Header.Get("Authorization")
	if jwtdata == "" {
		return JwtData{}, errors.New("token is invalid")
	}
	//jwtからユーザー情報を取得
	if !CheckJwt(jwtdata) {
		return JwtData{}, errors.New("token is invalid")
	}
	return UnpackJwt(jwtdata)
}
