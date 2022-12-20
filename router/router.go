package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	register(r, routes())
	return r
}

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

func register(r *gin.Engine, routes []route) {
	for _, route := range routes {
		r.Handle(route.method, route.path, route.handler)
	}
}

func routes() []route {
	return []route{
		{
			method:  http.MethodGet,
			path:    "/hello",
			handler: hello,
		},
	}
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
