package handler

import (
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary List openings
// @Description List openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	err := db.Find(&openings).Error

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "erros listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)

}
