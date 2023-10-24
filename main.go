package main

import (
	"github.com/Shiirookami/restapi_gin/controllers/productcontroller"
	"github.com/Shiirookami/restapi_gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	test := gin.Default()
	models.ConnectDB()

	test.GET("/api/products", productcontroller.Index)
	test.GET("/api/product/:id", productcontroller.Show)
	test.POST("/api/product", productcontroller.Create)
	test.PUT("/api/product/:id", productcontroller.Update)
	test.DELETE("/api/product", productcontroller.Delete)

	test.Run()
}
