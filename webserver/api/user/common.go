package user

import (
	"errors"
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

type NewPassword struct {
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
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
	if u, err := users.GetId(id); err != nil {
		return err
	} else {
		tmpuser := CreateUser{
			Name:      u.Name,
			Password:  user.OldPass,
			Authority: u.Authority,
		}
		if checkPassword(&tmpuser) {
			tmpuser.Password = user.NewPass
			if err := hashPassword(&tmpuser); err != nil {
				return err
			}
			if err := users.Update(id, tmpuser.Password); err != nil {
				return err
			}
		} else {
			return errors.New("password is not match")
		}
	}
	return nil
}
