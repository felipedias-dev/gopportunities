package handler

import (
	"net/http"

	"github.com/felipedias-dev/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningReq{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("update opening error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if request.IsEmpty() {
		logger.Error("validation error: at least one field must be valid")
		sendError(ctx, http.StatusBadRequest, "validation error: at least one field must be valid")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		logger.Error("update opening error: param id is missing")
		sendError(ctx, http.StatusBadRequest, "param id is missing")
		return
	}

	opening := schema.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Salary:   request.Salary,
		Link:     request.Link,
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("update opening error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}
