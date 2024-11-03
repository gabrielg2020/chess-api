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
			// case 3: // Get Bishop Move
			// 	getBishopMove(piece, row, col, chessboard)
			// case 4: // Get Rook Move
			// 	getRookMove(piece, row, col, chessboard)
			// case 5: // Get Queen Move
			// 	getQueenMove(piece, row, col, chessboard)
			// case 6: // Get King Move
			// 	getKingMove(piece, row, col, chessboard)
			default: // NOTE: Error on default when rest of cases are completed. for now add random move
				logger.Log.Debug("Default case hit. Adding random move")
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(0), HelperService.IntPtr(0),
					HelperService.IntPtr(7), HelperService.IntPtr(7),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				))
			}
		}
	}
	// c. For each piece, find all moves
	// 2. Filter for Legal Moves // TODO moveFilter for Legal Moves into separate function
	// a. Remove any move that goes off the board
	// b. Remove any move that place king in check
	// 3. Return Legal Moves
	// a. Return moves array
	return moves[0], nil
}

func getPawnMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	// NOTE: When calling any methods from `chessboard` that interact with board, we must flip toX and toY as fromX=col and fromY=row
	// methods such as: IsSquareEmpty, GetPiece, IsOpponent
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
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
					HelperService.IntPtr(toX), HelperService.IntPtr(toY),
					HelperService.IntPtr(promotionPiece*(-1*direction)),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				))
			}
		} else {
			// Create move
			logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
			moves = append(moves, entity.NewMoveEntity(
				HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
				HelperService.IntPtr(toX), HelperService.IntPtr(toY),
				HelperService.IntPtr(0),
				HelperService.BoolPtr(false), HelperService.BoolPtr(false),
				HelperService.IntPtr(0),
			))

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
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
					HelperService.IntPtr(toX), HelperService.IntPtr(toY),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(true),
					HelperService.IntPtr(0),
				))
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
					moves = append(moves, entity.NewMoveEntity(
						HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
						HelperService.IntPtr(toX), HelperService.IntPtr(toY),
						HelperService.IntPtr(promotionPiece*(-1*direction)),
						HelperService.BoolPtr(false), HelperService.BoolPtr(false),
						HelperService.IntPtr(pieceCaptured),
					))
				}
			} else {
				// Create move
				logger.Log.Debugf("getPawnMove: move added. moves array now contains %v move/s", len(moves))
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
					HelperService.IntPtr(toX), HelperService.IntPtr(toY),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(pieceCaptured),
				))
			}
		}
	}
	return moves, nil
}

func getKnightMove(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	// NOTE: When calling any methods from `chessboard` that interact with board, we must flip toX and toY as fromX=col and fromY=row
	// methods such as: IsSquareEmpty, GetPiece, IsOpponent
	var moves []entity.MoveEntityInterface

	deltaX := []int{1, 1, -1, -1, 2, 2, -2, -2}
	deltaY := []int{2, -2, 2, -2, 1, -1, 1, -1}

	for i := 0; i < 8; i++ {
		toX, toY := fromX+deltaX[i], fromY+deltaY[i]
		isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"fromX": fromX, "fromY": fromY,
				"toX": toX, "toY": toY,
			}).Error("failed checking square")
			return nil, errors.New("MoveService.getKnightMove: " + err.Error())
		}

		if isSquareEmpty {
			// Create move
			logger.Log.Debugf("getKnightMove: move added. moves array now contains %v move/s", len(moves))
			moves = append(moves, entity.NewMoveEntity(
				HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
				HelperService.IntPtr(toX), HelperService.IntPtr(toY),
				HelperService.IntPtr(0),
				HelperService.BoolPtr(false), HelperService.BoolPtr(false),
				HelperService.IntPtr(0),
			))
		} else {
			isOpponent, err := chessboard.IsOpponent(piece, toY, toX)
			if err != nil {
				logger.Log.WithFields(logrus.Fields{
					"fromX": fromX, "fromY": fromY,
					"toX": toX, "toY": toY,
				}).Error("failed checking square")
				return nil, errors.New("MoveService.getKnightMove: " + err.Error())
			}

			if isOpponent {
				// Create move
				pieceCaptured, err := chessboard.GetPiece(toY, toX)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"fromX": fromX, "fromY": fromY,
						"toX": toX, "toY": toY,
					}).Error("failed getting piece")
					return nil, errors.New("MoveService.getKnightMove: " + err.Error())
				}
				logger.Log.Debugf("getKnightMove: move added. moves array now contains %v move/s", len(moves))
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
					HelperService.IntPtr(toX), HelperService.IntPtr(toY),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(pieceCaptured),
				))
			}
		}
	}
	return moves, nil
}
