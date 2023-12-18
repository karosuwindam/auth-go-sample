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

func Init(pass string, r *gin.Engine, add ...any) {
	if pass == "" {
		pass = "/"
	}
	v1 := r.Group(pass)
	for _, a := range add {
		switch a.(type) {
		case func():
			a.(func())()
		}
	}
	for _, route := range routes {
		routePass(v1, route)
	}
	routes = []Route{}
}

func routePass(r *gin.RouterGroup, route Route) {
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

// ginによる許可確認
func CommonOption(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "X-POST, Content-Type, Authorization")
	c.Header("Access-Control-Max-Age", "86400")
	c.JSON(204, gin.H{
		"message": "ok",
	})
}
