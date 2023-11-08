package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("list openings error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusOK, openings)
}
