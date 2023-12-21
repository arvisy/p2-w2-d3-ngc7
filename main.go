package main

import (
	"ngc7/config"
	"ngc7/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := config.InitDB()

	// userRoutes := server.Group("/users")
	// {
	// 	userRoutes.POST("/register", handler.RegisterUser)
	// 	userRoutes.POST("/login", handler.LoginUser)
	// }

	// storeRoutes := server.Group("/stores")
	// {
	// 	storeRoutes.POST("/register", handler.RegisterStore)
	// 	storeRoutes.GET("", handler.GetStores)
	// 	storeRoutes.GET("/:id", handler.GetStoreByID)
	// 	storeRoutes.PUT("/:id", handler.UpdateStoreByID)
	// 	storeRoutes.DELETE("/:id", handler.DeleteStoreByID)
	// }

	productRoutes := server.Group("/products")
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
