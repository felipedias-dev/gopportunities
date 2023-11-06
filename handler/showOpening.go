package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		logger.Error("show opening error: param id is missing")
		sendError(ctx, http.StatusBadRequest, "param id is missing")
		return
	}

	opening := schema.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("show opening error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}
