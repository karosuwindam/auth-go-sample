package tables_test

import (
	"os"
	"suth-go-sample/config"
	"suth-go-sample/tables"
	"testing"
)

func TestTable(t *testing.T) {

	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_FILE", "test.db")
	config.Init()
	if err := tables.Init(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(config.DB.DBFile)

}
