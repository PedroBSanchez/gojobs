package handler

import (
	"net/http"

	"github.com/PedroBSanchez/gojobs.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	err := db.Find(&openings).Error

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "erros listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)

}
