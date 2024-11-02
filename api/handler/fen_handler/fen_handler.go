package FENHandler

import (
	"net/http"

	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gin-gonic/gin"
)

type FENHandler struct {
	fenService FENService.FENServiceInterface
}

func NewFENHandler(FENService FENService.FENServiceInterface) *FENHandler {
	return &FENHandler{fenService: FENService}
}

func (handler *FENHandler) ValidateFEN(ctx *gin.Context) {
	fen := ctx.Query("fen")

	err := handler.fenService.Validate(fen)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"valid":        err == nil,
			"errorMessage": "FENHandler.ValidateFEN: " + err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid": err == nil,
	})
}
