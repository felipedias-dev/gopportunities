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
	StatusCode int                    `json:"statusCode"`
	Data       schema.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	StatusCode int                    `json:"statusCode"`
	Data       schema.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	StatusCode int                    `json:"statusCode"`
	Data       schema.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	StatusCode int                      `json:"statusCode"`
	Data       []schema.OpeningResponse `json:"data"`
}
