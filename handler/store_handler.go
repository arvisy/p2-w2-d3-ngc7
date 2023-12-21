package handler

import (
	"errors"
	"net/http"
	"ngc7/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreHandler struct {
	DB *gorm.DB
}

func NewStoreHandler(db *gorm.DB) StoreHandler {
	return StoreHandler{DB: db}
}

func (s *StoreHandler) AddStore(ctx *gin.Context) {
	var store model.Store

	err := ctx.ShouldBindJSON(&store)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := s.DB.Create(&store)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, store)
}

func (s *StoreHandler) GetStores(ctx *gin.Context) {
	var stores []model.Store

	result := s.DB.Preload("Products").Find(&stores)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, stores)
}

func (s *StoreHandler) GetStoreByID(ctx *gin.Context) {
	storeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid store id",
		})
		return
	}

	var store model.Store
	result := s.DB.Preload("Products").First(&store, storeID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "store not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, store)
}

func (s *StoreHandler) DeleteStoreByID(ctx *gin.Context) {
	storeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid store id",
		})
		return
	}

	var store model.Store
	result := s.DB.First(&store, storeID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "store not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	result = s.DB.Delete(&store)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "store successfully deleted",
	})
}

func (s *StoreHandler) UpdateStoreByID(ctx *gin.Context) {
	storeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid store id",
		})
		return
	}

	var store model.Store
	result := s.DB.First(&store, storeID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "store not found",
		})
		return
	}

	var updateStore model.Store
	err = ctx.ShouldBind(&updateStore)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	store.Name = updateStore.Name

	result = s.DB.Save(&store)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, store)
}
