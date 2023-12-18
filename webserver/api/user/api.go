package user

import "suth-go-sample/webserver/api/common"

const API_NAME = "/user"

func Init() {
	if err := common.Add(API_NAME+"/:id", Read, common.GET); err != nil {
		panic(err)
	}
}
