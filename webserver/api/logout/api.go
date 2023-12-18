package logout

import "suth-go-sample/webserver/api/common"

const API_NAME = "/logout"

func Init() {
	if err := common.Add(API_NAME, LogoutPost, common.POST); err != nil {
		panic(err)
	}
}
