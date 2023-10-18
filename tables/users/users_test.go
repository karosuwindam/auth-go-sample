package users_test

import (
	"os"
	"suth-go-sample/config"
	"suth-go-sample/tables/users"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUsers(t *testing.T) {
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_FILE", "test.db")
	config.Init()
	db, err := gorm.Open(sqlite.Open(config.DB.DBFile), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(config.DB.DBFile)
	users.Init(db)
	if err := users.Add("test", "test", 1); err != nil {
		t.Fatal(err)
	}
	if err := users.Add("test", "test", 1); err == nil {
		t.Fatal("user already exists")
	}
	if user, err := users.Get("test"); err != nil {
		t.Fatal(err)
	} else {
		if user.Name != "test" {
			t.Fatal("user name is not test")
		}
		if user.Password != "test" {
			t.Fatal("user password is not test")
		}
		if user.Authority != 1 {
			t.Fatal("user authority is not 1")
		}
	}
	if err := users.Update(1, "test2"); err != nil {
		t.Fatal(err)
	}
	if user, err := users.Get("test"); err != nil {
		t.Fatal(err)
	} else {
		if user.Name != "test" {
			t.Fatal("user name is not test")
		}
		if user.Password != "test2" {
			t.Fatal("user password is not test2")
		}
		if user.Authority != 1 {
			t.Fatal("user authority is not 1")
		}
	}
	if err := users.Delete(1); err != nil {
		t.Fatal(err)
	}
	if _, err := users.Get("test"); err == nil {
		t.Fatal("user is not deleted")
	}

}
