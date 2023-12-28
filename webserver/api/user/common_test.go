package user

import (
	"fmt"
	"os"
	"suth-go-sample/config"
	"suth-go-sample/tables"
	"suth-go-sample/tables/users"
	"suth-go-sample/webserver/api/common"
	"testing"
)

// ユーザ操作のプライベート関数のテスト
func TestUserCommon(t *testing.T) {
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_FILE", "test.db")
	config.Init()
	tables.Init()

	defer os.Remove("test.db")
	name := "test"
	pass := "test"
	id := 1
	cUser := CreateUser{
		Name:      name,
		Password:  pass,
		Authority: common.ADMIN,
	}
	fmt.Println("test user common")
	t.Log("test hashPassword")
	if err := hashPassword(&cUser); err != nil {
		t.Error("hashPassword error")
	}
	fmt.Println("test HashPassword ok")
	if err := users.Add(cUser.Name, cUser.Password, cUser.Authority); err != nil {
		t.Error("users.Add error")
	}
	if !checkPassword(&CreateUser{Name: name, Password: pass}) {
		t.Error("checkPassword error")
	}
	if checkPassword(&CreateUser{Name: name, Password: "wrong"}) {
		t.Error("checkPassword error")
	}
	fmt.Println("test checkPassword ok")
	if list, err := makeUserList(); err != nil {
		t.Error("makeUserList error")
	} else {
		if len(list) != 1 {
			t.Error("makeUserList error")
		} else {
			if list[0].Name != name {
				t.Error("makeUserList error")
			} else {
				id = list[0].Id
			}
		}
	}
	fmt.Println("test makeUserList ok")
	nUserPass := NewPassword{
		OldPass: "newpass",
		NewPass: "newpass",
	}
	t.Log("test updatePassword")
	if err := updatePassword(id, &nUserPass); err == nil {
		t.Error("updatePassword error")
	}
	if !checkPassword(&CreateUser{Name: name, Password: pass}) {
		t.Error("updatePassword error")
	}
	t.Log("test updatePassword error ok")
	nUserPass.OldPass = pass
	if err := updatePassword(id, &nUserPass); err != nil {
		t.Error("updatePassword error")
	}
	t.Log("test updatePassword return ok")
	if !checkPassword(&CreateUser{Name: name, Password: nUserPass.NewPass}) {
		t.Error("updatePassword error")
	}
	t.Log("test updatePassword checkPassword ok")
	fmt.Println("test updatePassword ok")
	fmt.Println("test user common end")

}
