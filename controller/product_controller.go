package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	ProductUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUseCase: usecase,
	}
}

func (p *productController) GETProducts(ctx *gin.Context) {
	products, err := p.ProductUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}
