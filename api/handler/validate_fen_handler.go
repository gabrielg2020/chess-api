package handler

import (
	"net/http"

	"github.com/gabrielg2020/chess-api/api/service/validate"
	"github.com/gin-gonic/gin"
)

type ValidateFENHandler struct {
	validator validate.FENServiceInterface
}

func NewFENValidatorHandler(validator validate.FENServiceInterface) *ValidateFENHandler {
	return &ValidateFENHandler{validator: validator}
}

func (handler *ValidateFENHandler) ValidateFEN(ctx *gin.Context) {
	fen := ctx.Query("fen")

	isValid, err := handler.validator.Validate(fen)

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
