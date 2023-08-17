package handler

import (
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Update opening
// @Description Update a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body UpdateOpeningRequest true "Request body"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {

	openingId := ctx.Query("id")

	if openingId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	err = db.First(&opening, openingId).Error

	if err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	err = db.Save(&opening).Error

	if err != nil {
		logger.ErrorF("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)

}
