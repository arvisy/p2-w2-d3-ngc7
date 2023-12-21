package middleware

import (
	"net/http"
	"ngc7/helpers"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()

			switch e := err.Err.(type) {
			case *gin.Error:
				helpers.HandlerError(ctx, http.StatusBadRequest, "BAD_REQUEST", "Bad Request", e.Error())
			default:
				helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Internal Server Error", err.Error())
			}
		}
	}
}
