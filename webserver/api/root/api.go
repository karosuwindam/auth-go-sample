package root

import (
	"suth-go-sample/webserver/api/common"
)

func Init(pass string) {
	if err := common.Add(pass, Hello, common.ALL-common.OPTIONS); err != nil {
		panic(err)
	}
}
