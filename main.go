package main

import (
	categoryController "go/api_catalogue/controller/Category"
	productController "go/api_catalogue/controller/Product"
	userController "go/api_catalogue/controller/User"
	"go/api_catalogue/model"

	"go/api_catalogue/middleware"

	"github.com/gin-gonic/gin"
)
func main() {
	r:= gin.Default()
	model.ConnectDatabase()
	r.Static("uploads","./uploads")
	protectedRoutes := r.Group("/protected")
	api := r.Group("/api")
	{
		// Product routes
		api.GET("/products", productController.Index)
		api.GET("/product/:id", productController.Show)
		api.POST("/product", productController.Create)
		api.PUT("/product/:id", productController.Update)
		api.DELETE("/product/:id", productController.Delete)

		// Category routes
		api.GET("/categories", categoryController.Index)
		api.GET("/category/:id", categoryController.Show)
		protectedRoutes.Use(middleware.AuthenticationMiddleware())
		{
					protectedRoutes.POST("/category", categoryController.Create)
					protectedRoutes.PUT("/category/:id", categoryController.Update)
					protectedRoutes.DELETE("/category/:id", categoryController.Delete)

		}

		// User routes
		api.POST("/user/register", userController.Register)
		api.POST("/user/login", userController.Login)
	}
	r.Run()
}