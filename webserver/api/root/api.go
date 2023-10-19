package root

import (
	"suth-go-sample/webserver/api/common"
)

func Init() {
	if err := common.Add("/", Hello, common.ALL-common.OPTIONS); err != nil {
		panic(err)
	}
}
