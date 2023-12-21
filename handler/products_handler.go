package handler

import (
	"errors"
	"net/http"
	"ngc7/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) ProductHandler {
	return ProductHandler{DB: db}
}

func (p *ProductHandler) AddProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := p.DB.Create(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (p *ProductHandler) GetProduct(ctx *gin.Context) {
	var products []model.Product

	result := p.DB.Find(&products)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductHandler) GetProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	var product model.Product
	result := p.DB.Preload("Store").First(&product, productID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductHandler) DeleteProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	var product model.Product
	result := p.DB.First(&product, productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "product not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	result = p.DB.Delete(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product successfully deleted",
	})
}

func (p *ProductHandler) UpdateProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	var product model.Product
	result := p.DB.First(&product, productID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	var updateProduct model.Product
	err = ctx.ShouldBind(&updateProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product.Name = updateProduct.Name
	product.Description = updateProduct.Description
	product.ImageURL = updateProduct.ImageURL
	product.Price = updateProduct.Price
	product.StoreID = updateProduct.StoreID

	result = p.DB.Save(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, product)
}
