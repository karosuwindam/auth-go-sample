package register

import "suth-go-sample/webserver/api/common"

func Init() {
	if err := common.Add("/register", RegisterPut, common.PUT); err != nil {
		panic(err)
	}
	if err := common.Add("/register", UpdatePassword, common.POST); err != nil {
		panic(err)
	}
}
