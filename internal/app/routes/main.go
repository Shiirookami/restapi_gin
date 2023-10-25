package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Shiirookami/restapi_gin/controllers"
)

type Route struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}

func main() {
	r := gin.Default()

	routes := []Route{
		{
			Method:  "GET",
			Path:    "/products",
			Handler: .Index,
		},
		{
			Method:  "GET",
			Path:    "/product/:id",
			Handler: Show,
		},
		{
			Method:  "POST",
			Path:    "/product",
			Handler: Create,
		},
		{
			Method:  "PUT",
			Path:    "/product/:id",
			Handler: Update,
		},
		{
			Method:  "DELETE",
			Path:    "/product",
			Handler: Delete,
		},
	}
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	r.Run()
}
