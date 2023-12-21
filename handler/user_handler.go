package handler

import (
	"net/http"
	"ngc7/helpers"
	"ngc7/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{DB: db}
}

func (u *UserHandler) RegisterUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid request body", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to hash the password", err.Error())
		return
	}

	user.Password = string(hashedPassword)

	result := u.DB.Create(&user)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Error creating user", result.Error.Error())
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserHandler) LoginUser(ctx *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Invalid request body", err.Error())
		return
	}

	var user model.User
	result := u.DB.Where("username = ?", loginRequest.Username).First(&user)
	if result.Error != nil {
		helpers.HandlerError(ctx, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid username or password", "")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		helpers.HandlerError(ctx, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid username or password", "")
		return
	}

	token := "example_token"

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
