package MoveHandler

import (
	"github.com/gabrielg2020/chess-api/pkg/logger"
	"net/http"

	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gabrielg2020/chess-api/api/service/move_service"
	"github.com/gin-gonic/gin"
)

type MoveHandler struct {
	fenService  FENService.FENServiceInterface
	moveService MoveService.MoveServiceInterface
}

func NewMoveHandler(FENService FENService.FENServiceInterface, MoveService MoveService.MoveServiceInterface) *MoveHandler {
	return &MoveHandler{
		fenService:  FENService,
		moveService: MoveService,
	}
}

func (handler *MoveHandler) FindBestMove(ctx *gin.Context) {
	// Validate FEN
	fen := ctx.Query("fen")

	err := handler.fenService.Validate(fen)

	if err != nil {
		logger.Log.WithError(err).Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"valid":     err == nil,
			"error":     "inputted fen is invalid",
			"errorCode": http.StatusBadRequest,
		})
		return
	}

	// Parse FEN
	chessboard, err := handler.fenService.Parse(fen)

	if err != nil {
		logger.Log.WithError(err).Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "inputted fen failed to parsed",
			"errorCode":    http.StatusBadRequest,
			"chessboard":   chessboard,
		})
		return
	}

	// Find best move
	bestMove, boardErr := handler.moveService.FindBestMove(chessboard)

	if boardErr != nil {
		logger.Log.WithError(err).Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "failed to find best move",
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	chessNotation, chessNotationErr := bestMove.GetChessNotation()

	if chessNotationErr != nil {
		logger.Log.WithError(err).Error()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "failed to convert best move to chess notation",
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"move": chessNotation,
	})
}
