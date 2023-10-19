package register

import "suth-go-sample/webserver/api/common"

const API_NAME = "/register"

func Init() {
	if err := common.Add(API_NAME, RegisterPut, common.PUT); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME, UpdatePassword, common.POST); err != nil {
		panic(err)
	}
}
