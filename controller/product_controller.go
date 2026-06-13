package controller

import (
	"net/http"
	"products-api/model"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	//usecase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context){
	products := []model.Product{
		{
			ID: 1,
			Name: "Máquina de Lavar Louças",
			Price: 1750.00,
		},

	}
	
	ctx.JSON(http.StatusOK, products)
}