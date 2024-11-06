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
	FindPseudoLegalMoves(colour string, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error)
}

type MoveService struct{}

func NewMoveService() *MoveService {
	return &MoveService{}
}

func (service *MoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (entity.MoveEntityInterface, error) {
	// 1. Find Pseudo Legal Moves
	activeColour, err := chessboard.GetActiveColour()
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.FindBestMove:" + err.Error())
	}

	pseudoLegalMoves, err := service.FindPseudoLegalMoves(activeColour, chessboard)
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.FindBestMove:" + err.Error())
	}

	// 2. Filter for Legal Moves // TODO moveFilter for Legal Moves into separate function
	// a. Remove any move that goes off the board
	// b. Remove any move that place king in check
	// 3. Return Legal Moves
	// a. Return moves array
	logger.Log.Debugf("FindBestMove: moves array contains %v move/s", len(pseudoLegalMoves))
	return pseudoLegalMoves[0], nil
}

func (service *MoveService) FindPseudoLegalMoves(colour string, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	var moves []entity.MoveEntityInterface

	board, err := chessboard.GetBoard()
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
	}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == 0 || (colour == "w" && math.Signbit(float64(piece))) || (colour == "b" && !math.Signbit(float64(piece))) {
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
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, pawnMoves...)
			case 2: // Get Knight Move
				knightMoves, err := getKnightMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, knightMoves...)
			case 3: // Get Bishop Move
				bishopMoves, err := getBishopMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, bishopMoves...)
			case 4: // Get Rook Move
				rookMoves, err := getRookMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, rookMoves...)
			case 5: // Get Queen Move
				queenMoves, err := getQueenMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, queenMoves...)
			case 6: // Get King Move
				kingMoves, err := getKingMove(piece, row, col, chessboard)
				if err != nil {
					logger.Log.WithFields(logrus.Fields{
						"board": board,
						"row":   row, "col": col,
					}).Error()
					return nil, errors.New("MoveService.findPseudoLegalMoves:" + err.Error())
				}
				moves = append(moves, kingMoves...)
			default: // Error if piece not found
				logger.Log.WithFields(logrus.Fields{
					"board": board,
					"row":   row, "col": col,
				}).Error()
				return nil, errors.New("MoveService.findPseudoLegalMoves: piece not found")
			}
		}
	}

	return moves, nil
}

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

func getCastlingMoves(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	var moves []entity.MoveEntityInterface

	// Get castling rights
	castlingRights, err := chessboard.GetCastlingRights()
	if err != nil {
		logger.Log.Error()
		return nil, errors.New("MoveService.getCastlingMoves: " + err.Error())
	}

	var kingSideRight, queenSideRight rune
	var kingSideX, queenSideX int

	if piece > 0 {
		kingSideRight, queenSideRight = 'K', 'Q'
		kingSideX, queenSideX = fromX+2, fromX-2
	} else {
		kingSideRight, queenSideRight = 'k', 'q'
		kingSideX, queenSideX = fromX-2, fromX+3
	}

	// King Side Castling
	if strings.ContainsRune(castlingRights, kingSideRight) {
		canCastle, err := canCastleKingSide(piece, fromY, fromX, chessboard)
		if err != nil {
			logger.Log.Error()
			return nil, errors.New("MoveService.getCastlingMoves: " + err.Error())
		}
		if canCastle {
			addMove(fromX, fromY, kingSideX, fromY, 0, true, false, 0, &moves)
		}
	}

	// Queen Side Castling
	if strings.ContainsRune(castlingRights, queenSideRight) {
		canCastle, err := canCastleQueenSide(piece, fromY, fromX, chessboard)
		if err != nil {
			logger.Log.Error()
			return nil, errors.New("MoveService.getCastlingMoves: " + err.Error())
		}
		if canCastle {
			addMove(fromX, fromY, queenSideX, fromY, 0, true, false, 0, &moves)
		}
	}

	return moves, nil
}

func canCastleKingSide(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) (bool, error) {
	pathX := []int{fromX + 1, fromX + 2}

	for _, x := range pathX {
		isSquareEmpty, err := chessboard.IsSquareEmpty(fromY, x)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"fromX": fromX, "fromY": fromY,
				"x": x,
			}).Error("failed checking square")
			return false, errors.New("MoveService.canCastleKingSide: " + err.Error())
		}
		if !isSquareEmpty {
			return false, nil
		}
	}

	rookX := fromX + 3
	rookPiece, err := chessboard.GetPiece(fromY, rookX)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"fromX": fromX, "fromY": fromY,
			"rookX": rookX,
		}).Error("failed getting rook piece")
		return false, errors.New("MoveService.canCastleKingSide: " + err.Error())
	}
	expectedRook := 4
	if piece < 0 {
		expectedRook = -4
	}
	if rookPiece != expectedRook {
		return false, nil
	}

	return true, nil
}

func canCastleQueenSide(piece int, fromY int, fromX int, chessboard entity.ChessboardEntityInterface) (bool, error) {
	pathX := []int{fromX - 1, fromX - 2, fromX - 3}

	for _, x := range pathX {
		isSquareEmpty, err := chessboard.IsSquareEmpty(fromY, x)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"fromX": fromX, "fromY": fromY,
				"x": x,
			}).Error("failed checking square")
			return false, errors.New("MoveService.canCastleKingSide: " + err.Error())
		}
		if !isSquareEmpty {
			return false, nil
		}
	}

	rookX := fromX - 4
	rookPiece, err := chessboard.GetPiece(fromY, rookX)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"fromX": fromX, "fromY": fromY,
			"rookX": rookX,
		}).Error("failed getting rook piece")
		return false, errors.New("MoveService.canCastleKingSide: " + err.Error())
	}
	expectedRook := 4
	if piece < 0 {
		expectedRook = -4
	}
	if rookPiece != expectedRook {
		return false, nil
	}

	return true, nil
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

// TODO castling logic is missing
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
