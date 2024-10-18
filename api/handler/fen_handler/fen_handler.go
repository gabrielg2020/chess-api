package FENHandler

import (
	"net/http"

	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gin-gonic/gin"
)

type FENHandler struct {
	service FENService.FENServiceInterface
}

func NewFENHandler(service FENService.FENServiceInterface) *FENHandler {
	return &FENHandler{service: service}
}

func (handler *FENHandler) ValidateFEN(ctx *gin.Context) {
	fen := ctx.Query("fen")

	isValid, err := handler.service.Validate(fen)

	if  err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"valid":        isValid,
			"errorMessage": err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"valid": isValid,
	})
}
