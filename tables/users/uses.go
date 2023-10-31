package users

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string // ユーザー名
	Password  string // パスワード
	Authority int    // 権限
}

var db *gorm.DB

func Init(tmp *gorm.DB) error {
	db = tmp
	return db.AutoMigrate(&User{})
}

func Add(Name string, Password string, Authority int) error {
	users := []User{}
	if results := db.Where("name = ?", Name).Find(&users); results.Error != nil {
		return results.Error
	} else {
		if len(users) > 0 {
			return errors.New("user already exists")
		}
	}
	user := User{
		Name:      Name,
		Password:  Password,
		Authority: Authority,
	}
	if results := db.Create(&user); results.Error != nil {
		return results.Error
	}
	return nil
}

func Delete(Id int) error {
	if results := db.Delete(&User{}, Id); results.Error != nil {
		return results.Error
	}
	return nil
}

func Update(Id int, Password string) error {
	user := User{}
	if results := db.First(&user, Id); results.Error != nil {
		return results.Error
	}
	user.Password = Password
	db.Save(&user)
	return nil
}

func Get(Name string) (User, error) {
	user := User{}
	if results := db.Where("name = ?", Name).First(&user); results.Error != nil {
		return user, results.Error
	}
	return user, nil
}

func (u *User) ReadAuth() string {
	switch u.Authority {
	case 3:
		return "admin"
	case 2:
		return "admin"
	case 1:
		return "user"
	case 0:
		return "guest"
	default:
		return "unknown"
	}
}
