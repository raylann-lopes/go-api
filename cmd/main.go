package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.DbConnection()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", productController.GETProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductsById)

	server.Run(":8000")
}
