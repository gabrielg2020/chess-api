package BestMoveHandler

import (
	"net/http"

	"github.com/gabrielg2020/chess-api/api/service/fen_service"
	"github.com/gabrielg2020/chess-api/api/service/best_move_service"
	"github.com/gin-gonic/gin"
)

type BestMoveHandler struct {
	fenService FENService.FENServiceInterface
	bestMoveService BestMoveService.BestMoveServiceInterface
}

func NewBestMoveHandler(FENService FENService.FENServiceInterface, BestMoveService BestMoveService.BestMoveServiceInterface) *BestMoveHandler {
	return &BestMoveHandler{
		fenService: FENService,
		bestMoveService: BestMoveService,
	}
}

func (handler *BestMoveHandler) FindBestMove(ctx *gin.Context) {
	// Validate FEN
	fen := ctx.Query("fen")

	err := handler.fenService.Validate(fen)

	if  err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"valid":        (err == nil),
			"errorMessage": err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	// Parse FEN
	chessboard , err := handler.fenService.Parse(fen)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	// Find best move
	bestmove, err := handler.bestMoveService.FindBestMove(chessboard)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bestMove": bestmove,
	})
}

