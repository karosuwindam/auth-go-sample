package user

import "suth-go-sample/webserver/api/common"

const API_NAME = "/user"

func Init() {
	if err := common.Add(API_NAME, Create, common.PUT); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME, Update, common.POST); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME+"/:id", Read, common.GET); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME+"/:id", Delete, common.DELETE); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME+"/:id", UpdateById, common.POST); err != nil {
		panic(err)
	}
	if err := common.Add(API_NAME+"/list", ReadList, common.GET); err != nil {
		panic(err)
	}
}
