package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening/:id [delete]
func ExcludeOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		logger.Error("exclude opening error: param id is missing")
		sendError(ctx, http.StatusBadRequest, "param id is missing")
		return
	}

	opening := schema.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("exclude opening error: %v", err.Error())
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("exclude opening error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusNoContent, nil)
}
