package login

import "suth-go-sample/webserver/api/common"

func Init() {
	if err := common.Add("/login", LoginPost, common.POST); err != nil {
		panic(err)
	}
	if err := common.Add("/login", LoginGet, common.GET); err != nil {
		panic(err)
	}
}
