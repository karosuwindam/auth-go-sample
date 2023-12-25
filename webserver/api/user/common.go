package user

import (
	"suth-go-sample/config"
	"suth-go-sample/tables/users"

	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Auth int    `json:"auth"`
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

// データベースからユーザ情報をのリストを作成する
func makeUserList() ([]User, error) {
	var userList []User
	if users, err := users.GetAll(); err != nil {
		return nil, err
	} else {
		for i := 0; i < len(users); i++ {
			userList = append(userList, User{
				Id:   int(users[i].Model.ID),
				Name: users[i].Name,
				Auth: users[i].Authority,
			})
		}
	}
	return userList, nil
}

// idとユーザと新旧パスワードを指定してパスワードを更新する
func updatePassword(id int, user *NewPassword) error {
	//ToDo: パスワードの更新処理を実装する
	return nil
}
