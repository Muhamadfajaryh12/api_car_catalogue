package main

import (
	categoryController "go/api_catalogue/controller/Category"
	productController "go/api_catalogue/controller/Product"
	"go/api_catalogue/model"

	"github.com/gin-gonic/gin"
)
func main() {
	r:= gin.Default()
	model.ConnectDatabase()

	r.GET("/api/products",productController.Index)
	r.GET("/api/product/:id",productController.Show)
	r.POST("/api/product",productController.Create)
	r.PUT("/api/product/:id",productController.Update)
	r.DELETE("/api/product/:id",productController.Delete)

	
	r.GET("/api/categorys",categoryController.Index)
	r.GET("/api/category/:id",categoryController.Show)
	r.POST("/api/category",categoryController.Create)
	r.PUT("/api/category/:id",categoryController.Update)
	r.DELETE("/api/category/:id",categoryController.Delete)

	r.Run()
}