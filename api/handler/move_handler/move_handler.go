package MoveHandler

import (
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"valid":        err == nil,
			"errorMessage": "MoveHandler.FindBestMove" + err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	// Parse FEN
	chessboard, err := handler.fenService.Parse(fen)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "MoveHandler.FindBestMove" + err.Error(),
			"errorCode":    http.StatusBadRequest,
			"chessboard":   chessboard,
		})
		return
	}

	// Find best move
	bestMove, boardErr := handler.moveService.FindBestMove(chessboard)
	chessNotation, chessNotationErr := bestMove.GetChessNotation()

	if boardErr != nil || chessNotationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "MoveHandler.FindBestMove" + err.Error(),
			"errorCode":    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"move": chessNotation,
	})
}
