package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/:id [get]
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
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}
