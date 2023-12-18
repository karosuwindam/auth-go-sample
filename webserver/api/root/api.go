package root

import (
	"suth-go-sample/webserver/api/common"
)

const API_NAME = "/"

func Init() {
	if err := common.Add(API_NAME, Hello, common.ALL-common.OPTIONS); err != nil {
		panic(err)
	}
}
