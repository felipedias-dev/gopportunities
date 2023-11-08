package handler

import (
	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, statusCode int, msg string) {
	ctx.JSON(statusCode, gin.H{
		"errorCode": statusCode,
		"message":   msg,
	})
}

func sendSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"data":       data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                 `json:"message"`
	Data    schema.OpeningResponse `json:"data"`
}
