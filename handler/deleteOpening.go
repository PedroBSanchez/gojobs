package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	openingId := ctx.Query("id")

	if openingId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}

	opening := schemas.Opening{}

	//Find Opening
	err := db.First(&opening, openingId).Error

	if err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", openingId))
		return
	}

	//Delete Opening
	err = db.Delete(&opening).Error

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", openingId))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)

}
