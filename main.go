package main

import (
	"ngc7/config"
	"ngc7/handler"
	"ngc7/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := config.InitDB()

	userRoutes := server.Group("/users")
	userRoutes.Use(middleware.ErrorHandler())
	userHandler := handler.NewUserHandler(db)
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.LoginUser)
	}

	storeRoutes := server.Group("/stores")
	storeRoutes.Use(middleware.ErrorHandler())
	storeHandler := handler.NewStoreHandler(db)
	{
		storeRoutes.POST("/register", storeHandler.AddStore)
		storeRoutes.GET("", storeHandler.GetStores)
		storeRoutes.GET("/:id", storeHandler.GetStoreByID)
		storeRoutes.PUT("/:id", storeHandler.UpdateStoreByID)
		storeRoutes.DELETE("/:id", storeHandler.DeleteStoreByID)
	}

	productRoutes := server.Group("/products")
	productRoutes.Use(middleware.ErrorHandler())
	// productRoutes.Use(middleware.AuthMiddleware)
	productHandler := handler.NewProductHandler(db)
	{
		productRoutes.POST("", productHandler.AddProduct)
		productRoutes.GET("", productHandler.GetProduct)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProductByID)
		productRoutes.DELETE("/:id", productHandler.DeleteProductByID)
	}

	server.Run(":8080")
}
