package root

import (
	"suth-go-sample/webserver/api/common"
)

func Init() {
	if err := common.Add("/", Hello, common.ALL); err != nil {
		panic(err)
	}
}
