package handler

import (
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error
	if err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)

}
