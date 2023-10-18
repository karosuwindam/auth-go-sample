package tables

import (
	"errors"
	"suth-go-sample/config"
	"suth-go-sample/tables/users"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type funcList struct {
	Func func(*gorm.DB) error
}

var listfunc []funcList = []funcList{
	{users.Init},
}

func Init() error {
	var err error
	if config.DB.DBType == "sqlite" {
		db, err = gorm.Open(sqlite.Open(config.DB.DBFile), &gorm.Config{})
	} else {
		return errors.New("DBType is not sqlite")
	}
	if err != nil {
		return err
	}
	for _, f := range listfunc {
		if err = f.Func(db); err != nil {
			return err
		}
	}
	return nil
}
