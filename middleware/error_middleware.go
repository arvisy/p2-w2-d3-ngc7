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

			if e, ok := err.Err.(*helpers.ErrorResponse); ok {
				ctx.JSON(e.StatusCode, e)
			} else {
				helpers.HandlerError(ctx, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Internal server error", "")
			}
		}
	}
}
