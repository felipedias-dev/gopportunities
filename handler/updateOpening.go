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

	opening := schema.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("update opening error: %v", err.Error())
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Remote != opening.Remote {
		opening.Remote = request.Remote
	}
	if request.Salary != opening.Salary && request.Salary != 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("update opening error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "update opening error")
		return
	}

	sendSuccess(ctx, http.StatusOK, opening)
}
