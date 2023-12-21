package handler

import (
	"errors"
	"net/http"
	"ngc7/helpers"
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
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid request body", err.Error())
		return
	}

	result := p.DB.Create(&product)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error creating product", result.Error.Error())
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (p *ProductHandler) GetProduct(ctx *gin.Context) {
	var products []model.Product

	result := p.DB.Find(&products)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error retrieving products", result.Error.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductHandler) GetProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid product ID", err.Error())
		return
	}

	var product model.Product
	result := p.DB.Preload("Store").First(&product, productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			helpers.HandlerError(ctx, http.StatusNotFound, "NOT_FOUND", "Product not found", "")
			return
		}

		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error retrieving product", result.Error.Error())
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductHandler) DeleteProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid product ID", err.Error())
		return
	}

	var product model.Product
	result := p.DB.First(&product, productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			helpers.HandlerError(ctx, http.StatusNotFound, "NOT_FOUND", "Product not found", "")
			return
		}

		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error retrieving product", result.Error.Error())
		return
	}

	result = p.DB.Delete(&product)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error deleting product", result.Error.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product successfully deleted",
	})
}

func (p *ProductHandler) UpdateProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid product ID", err.Error())
		return
	}

	var product model.Product
	result := p.DB.First(&product, productID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			helpers.HandlerError(ctx, http.StatusNotFound, "NOT_FOUND", "Product not found", "")
			return
		}

		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error retrieving product", result.Error.Error())
		return
	}

	var updateProduct model.Product
	err = ctx.ShouldBind(&updateProduct)
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid request body", err.Error())
		return
	}

	product.Name = updateProduct.Name
	product.Description = updateProduct.Description
	product.ImageURL = updateProduct.ImageURL
	product.Price = updateProduct.Price
	product.StoreID = updateProduct.StoreID

	result = p.DB.Save(&product)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error updating product", result.Error.Error())
		return
	}

	ctx.JSON(http.StatusOK, product)
}
