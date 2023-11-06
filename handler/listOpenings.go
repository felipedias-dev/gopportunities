package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("list openings error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusOK, openings)
}
