package login

import "suth-go-sample/webserver/api/common"

const API_NAME = "/login"

func Init() {
	if err := common.Add(API_NAME, LoginPost, common.POST); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME, LoginGet, common.GET); err != nil {
		panic(err)
	}
}
