package handler

import (
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Show opening
// @Description Show a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {

	openingId := ctx.Query("id")

	if openingId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	opening := schemas.Opening{}

	err := db.First(&opening, openingId).Error

	if err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)

}
