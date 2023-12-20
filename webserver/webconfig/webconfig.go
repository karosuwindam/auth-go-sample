package webconfig

import (
	"errors"
	"fmt"
	"suth-go-sample/config"
	"suth-go-sample/tables/users"

	"golang.org/x/crypto/bcrypt"
)

func Init() error {
	if err := dbInit(); err != nil {
		return err
	}
	return nil
}

func dbInit() error {
	if us, err := users.GetAll(); err == nil {
		if len(us) != 0 {
			for _, u := range us {
				if u.Authority >= 2 {
					return nil
				}
			}
			return errors.New("admin user not found")
		} else { // ユーザがいない場合はadmin初期ユーザを作成
			pass, err := hashPassword("admin", "admin")
			if err != nil {
				return err
			}
			if err := users.Add("admin", pass, 3); err != nil {
				return err
			}
			fmt.Println("admin user created\n id: admin\n pass: admin")
		}
	} else {
		return errors.New("db init error")
	}
	return nil
}

// パスワードをペッパー文字列でハッシュ化する
func hashPassword(name, password string) (string, error) {
	if name == "" || password == "" {
		return "", errors.New("invalid name or password")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(name+password+config.JWT.Pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}
