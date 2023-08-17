package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

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
