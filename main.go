package main

import {
	"github.com/gin-gonic/gin"
	"github.com/Shiirookami/restapi_gin/models"
	"github.com/Shiirookami/restapi_gin/controllers/ProductController"

}

func main() {
	test := gin.Default()
	models.ConnectDB()

	test.Get("/api/products", ProductController.Index)
	test.Get("/api/products/:id", ProductController.Show)
	test.Post("/api/products", ProductController.Create)
	test.Put("/api/products/:id", ProductController.Update)
	test.Delete("/api/products", ProductController.Delete)

	test.Run()
}
