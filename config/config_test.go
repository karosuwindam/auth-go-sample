package config_test

import (
	"os"
	"suth-go-sample/config"
	"testing"
)

func TestConfig(t *testing.T) {
	os.Setenv("WEB_HOST", "localhost")
	os.Setenv("WEB_PORT", "8080")
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_FILE", "test.db")
	config.Init()
	if config.Web.Host != "localhost" {
		t.Errorf("Web.Host is not localhost")
	}
	if config.Web.Port != "8080" {
		t.Errorf("Web.Port is not 8080")
	}
	if config.DB.DBType != "sqlite" {
		t.Errorf("DB.DBType is not sqlite")
	}
	if config.DB.DBFile != "test.db" {
		t.Errorf("DB.DBFile is not test.db")
	}

}
