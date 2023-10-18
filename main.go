package main

import (
	"suth-go-sample/config"
	"suth-go-sample/tables"
	"suth-go-sample/webserver"
)

func main() {
	config.Init()
	tables.Init()
	webserver.Init()
	webserver.Start()
}
