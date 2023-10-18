package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	GET = 1 << iota
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD
)

const (
	ALL = GET | POST | PUT | DELETE | PATCH | OPTIONS | HEAD
)

const (
	GEST = iota
	USER
	ADMIN
)

type Route struct {
	Pass    string
	Handler func(c *gin.Context)
	Method  int
}

var routes = []Route{}

func Add(pass string, handler func(c *gin.Context), method int) error {
	for _, route := range routes {
		if route.Pass == pass && route.Method == method {
			return fmt.Errorf("route already exists")
		}
	}
	routes = append(routes, Route{pass, handler, method})
	return nil
}

func Init(pass string, r *gin.Engine) {
	if pass == "/" {
		pass = ""
	}
	for _, route := range routes {
		if route.Pass[:1] != "/" {
			route.Pass = "/" + route.Pass
		}
		route.Pass = pass + route.Pass
		routePass(r, route)
	}
	routes = []Route{}
}

func routePass(r *gin.Engine, route Route) {
	if route.Method&GET != 0 {
		r.GET(route.Pass, route.Handler)
	}
	if route.Method&POST != 0 {
		r.POST(route.Pass, route.Handler)
	}
	if route.Method&PUT != 0 {
		r.PUT(route.Pass, route.Handler)
	}
	if route.Method&DELETE != 0 {
		r.DELETE(route.Pass, route.Handler)
	}
	if route.Method&PATCH != 0 {
		r.PATCH(route.Pass, route.Handler)
	}
	if route.Method&OPTIONS != 0 {
		r.OPTIONS(route.Pass, route.Handler)
	}
	if route.Method&HEAD != 0 {
		r.HEAD(route.Pass, route.Handler)
	}
}
