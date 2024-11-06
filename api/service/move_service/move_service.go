package MoveService

import (
	"errors"
	"github.com/gabrielg2020/chess-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"math"

	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/gabrielg2020/chess-api/api/service/helper_service"
)

type MoveServiceInterface interface {
	FindBestMove(chessboard entity.ChessboardEntityInterface) (entity.MoveEntityInterface, error)
}

type MoveService struct{}

func NewMoveService() *MoveService {
	return &MoveService{}
}

func (service *MoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (entity.MoveEntityInterface, error) {
	// 1. Find Pseudo Legal Moves // TODO move Pseudo Legal Moves into separate function
	// a. Create a moves array
	var moves []entity.MoveEntityInterface
	board, err := chessboard.GetBoard()
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.FindBestMove:" + err.Error())
	}
	activeColour, err := chessboard.GetActiveColour()
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.FindBestMove:" + err.Error())
	}
	// b. Loop through the board
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == 0 || (activeColour == "w" && math.Signbit(float64(piece))) || (activeColour == "b" && !math.Signbit(float64(piece))) {
				continue
			}
			switch math.Abs(float64(piece)) {
			case 1: // Get Pawn Move
				pawnMoves, err := getPawnMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, pawnMoves...)
			case 2: // Get Knight Move
				knightMoves, err := getKnightMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, knightMoves...)
			case 3: // Get Bishop Move
				bishopMoves, err := getBishopMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, bishopMoves...)
			case 4: // Get Rook Move
				rookMoves, err := getRookMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, rookMoves...)
			case 5: // Get Queen Move
				queenMoves, err := getQueenMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, queenMoves...)
			case 6: // Get King Move
				kingMoves, err := getKingMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.FindBestMove:" + err.Error())
				}
				moves = append(moves, kingMoves...)
			default: // Error if piece not found
				logger.Log.WithFields(logrus.Fields{
					"board": board,
					"row":   row, "col": col,
				}).Error()
				return nil, errors.New("MoveService.FindBestMove: piece not found")
			}
		}
	}
	// c. For each piece, find all moves
	// 2. Filter for Legal Moves // TODO moveFilter for Legal Moves into separate function
	// a. Remove any move that goes off the board
	// b. Remove any move that place king in check
	// 3. Return Legal Moves
	// a. Return moves array
	logger.Log.Debugf("FindBestMove: moves array contains %v move/s", len(moves))
	return moves[0], nil
}

// TODO needs testing
func generateMoves(piece int, fromY int, fromX int, deltaXs []int, deltaYs []int, isSliding bool, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	var moves []entity.MoveEntityInterface

	for i := 0; i < len(deltaXs); i++ {
		deltaX, deltaY := deltaXs[i], deltaYs[i]
		toX, toY := fromX+deltaX, fromY+deltaY

		if isSliding {
			for chessboard.IsWithinBounds(toY, toX) {
				moveAdded, err := tryAddMove(piece, fromY, fromX, toY, toX, &moves, chessboard)
				if err != nil {
					logger.Log.Error()
					return nil, errors.New("MoveService.generateMoves: " + err.Error())
				}
				if !moveAdded {
					break
				}
				toX += deltaX
				toY += deltaY
			}
		} else {
			if chessboard.IsWithinBounds(toY, toX) {
				_, err := tryAddMove(piece, fromY, fromX, toY, toX, &moves, chessboard)
				if err != nil {
					logger.Log.Error()
					return nil, errors.New("MoveService.generateMoves: " + err.Error())
				}
			}
		}
	}
	return moves, nil
}

func tryAddMove(piece int, fromY int, fromX int, toY int, toX int, moves *[]entity.MoveEntityInterface, chessboard entity.ChessboardEntityInterface) (bool, error) {
	isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"fromX": fromX, "fromY": fromY,
			"toX": toX, "toY": toY,
		}).Error("failed checking square")
		return false, errors.New("MoveService.tryAddMove: " + err.Error())
	}

	if isSquareEmpty {
		addMove(fromX, fromY, toX, toY, 0, false, false, 0, moves)
		return true, nil
	} else {
		isOpponent, err := chessboard.IsOpponent(piece, toY, toX)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"fromX": fromX, "fromY": fromY,
				"toX": toX, "toY": toY,
			}).Error("failed checking is opponent")
			return false, errors.New("MoveService.tryAddMove: " + err.Error())
		}

		if isOpponent {
			pieceCaptured, err := chessboard.GetPiece(toY, toX)
			if err != nil {
				logger.Log.WithFields(logrus.Fields{
					"fromX": fromX, "fromY": fromY,
					"toX": toX, "toY": toY,
				}).Error("failed getting piece")
				return false, errors.New("MoveService.tryAddMove: " + err.Error())
			}
			addMove(fromX, fromY, toX, toY, 0, false, false, pieceCaptured, moves)
			return true, nil // Stop checking after capturing a piece
		}
		return false, nil // Blocked by own piece
	}
}

func addMove(fromX int, fromY int, toX int, toY int, promotionPiece int, isCastling bool, isEnPassant bool, pieceCaptured int, moves *[]entity.MoveEntityInterface) {
	*moves = append(*moves, entity.NewMoveEntity(
		HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
		HelperService.IntPtr(toX), HelperService.IntPtr(toY),
		HelperService.IntPtr(promotionPiece),
		HelperService.BoolPtr(isCastling), HelperService.BoolPtr(isEnPassant),
		HelperService.IntPtr(pieceCaptured),
	))
}

// NOTE: When calling any methods from `chessboard` that interact with board, we must flip toX and toY as fromX=col and fromY=row
// methods such as: IsSquareEmpty, GetPiece, IsOpponent

func getPawnMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	var moves []entity.MoveEntityInterface

	// Find startRank, promotionRank and direction
	var direction, startRank, promotionRank int
	if piece > 0 { // White
		direction = -1
		startRank = 6
		promotionRank = 0
	} else { // Black
		direction = 1
		startRank = 1
		promotionRank = 7
	}

	// Move 1 forward
	toX, toY := fromX, fromY+direction
	isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"fromX": fromX, "fromY": fromY,
			"toX": toX, "toY": toY,
		}).Error("failed checking 1 square forward")
		return nil, errors.New("MoveService.getPawnMove: " + err.Error())
	}

	if isSquareEmpty {
		if toY == promotionRank {
			for _, promotionPiece := range []int{2, 3, 4, 5} { // Create a move for each piece it can promote too
				logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
				addMove(fromX, fromY, toX, toY, promotionPiece*(-1*direction), false, false, 0, &moves)
			}
		} else {
			// Create move
			logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
			addMove(fromX, fromY, toX, toY, 0, false, false, 0, &moves)

			// Move 2 forward
			toY += direction

			isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)

			if err != nil {
				logger.Log.WithFields(logrus.Fields{
					"fromX": fromX, "fromY": fromY,
					"toX": toX, "toY": toY,
				}).Error("failed checking 2 squares forward")
				return nil, errors.New("MoveService.getPawnMove: " + err.Error())
			}
			if fromY == startRank && isSquareEmpty {
				// Don't check if can promote because a pawn can never promote off first move
				// Create move
				logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
				addMove(fromX, fromY, toX, toY, 0, false, true, 0, &moves)
			}
		}
	}
	// Check diagonal moves
	for _, deltaX := range []int{-1, 1} {
		toX, toY := fromX+deltaX, fromY+direction
		isOpponent, err := chessboard.IsOpponent(piece, toY, toX)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"fromX": fromX, "fromY": fromY,
				"toX": toX, "toY": toY,
				"deltaX": deltaX,
			}).Errorf("failed checking %v", deltaX)
			return nil, errors.New("MoveService.getPawnMove: " + err.Error())
		}

		if isOpponent {
			pieceCaptured, err := chessboard.GetPiece(toY, toX)

			if err != nil {
				logger.Log.WithFields(logrus.Fields{
					"fromX": fromX, "fromY": fromY,
					"toX": toX, "toY": toY,
					"deltaX": deltaX,
				}).Errorf("failed getting piece")
				return nil, errors.New("MoveService.getPawnMove: " + err.Error())
			}

			if pieceCaptured == 0 { // En Passant capture
				pieceCaptured = piece * -1
			}

			if toY == promotionRank {
				for _, promotionPiece := range []int{2, 3, 4, 5} { // Create a move for each piece it can promote too
					logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
					addMove(fromX, fromY, toX, toY, promotionPiece*(-1*direction), false, false, pieceCaptured, &moves)
				}
			} else {
				// Create move
				logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
				addMove(fromX, fromY, toX, toY, 0, false, false, pieceCaptured, &moves)
			}
		}
	}
	return moves, nil
}

// FEAT can create a move gen with delta's as input

func getKnightMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	deltaXs := []int{1, 1, -1, -1, 2, 2, -2, -2}
	deltaYs := []int{2, -2, 2, -2, 1, -1, 1, -1}
	return generateMoves(piece, fromY, fromX, deltaXs, deltaYs, false, chessboard)
}

func getBishopMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	deltaXs := []int{1, -1, 1, -1}
	deltaYs := []int{1, 1, -1, -1}
	return generateMoves(piece, fromY, fromX, deltaXs, deltaYs, true, chessboard)
}

func getRookMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	deltaXs := []int{1, -1, 0, 0}
	deltaYs := []int{0, 0, 1, -1}
	return generateMoves(piece, fromY, fromX, deltaXs, deltaYs, true, chessboard)
}

func getQueenMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	deltaXs := []int{1, -1, 1, -1, 1, -1, 0, 0}
	deltaYs := []int{1, 1, -1, -1, 0, 0, 1, -1}
	return generateMoves(piece, fromY, fromX, deltaXs, deltaYs, true, chessboard)
}

func getKingMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	deltaXs := []int{1, 1, 1, 0, 0, -1, -1, -1}
	deltaYs := []int{1, 0, -1, 1, -1, 1, 0, -1}
	return generateMoves(piece, fromY, fromX, deltaXs, deltaYs, false, chessboard)
}
