package handler

import (
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
