package helpers

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(statusCode int, errorCode, message, detail string) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
		Detail:     detail,
	}
}

func HandlerError(ctx *gin.Context, statusCode int, errorCode, message, detail string) {
	response := NewErrorResponse(statusCode, errorCode, message, detail)
	ctx.JSON(statusCode, response)
	ctx.Abort()
}
