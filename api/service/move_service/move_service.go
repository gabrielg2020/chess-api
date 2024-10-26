package MoveService

import (
	"errors"
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
		return nil, errors.New("failed to retrieve chessboard")
	}
		// b. Loop through the board
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == 0 {
				continue
			}
			switch math.Abs(float64(piece)) {
			case 1: // Get Pawn Move
				pawnMoves, err := getPawnMove(piece, col, row, chessboard)
				if err != nil {
					return nil, errors.New("failed to get pawn moves")
				}
				moves = append(moves, pawnMoves...)
			// case 2: // Get Knight Move
			// 	getKnightMove(piece, row, col, chessboard)
			// case 3: // Get Bishop Move
			// 	getBishopMove(piece, row, col, chessboard)
			// case 4: // Get Rook Move
			// 	getRookMove(piece, row, col, chessboard)
			// case 5: // Get Queen Move
			// 	getQueenMove(piece, row, col, chessboard)
			// case 6: // Get King Move
			// 	getKingMove(piece, row, col, chessboard)
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

// TODO needs to be tested ... :(
func getPawnMove(piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
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

	// 1 move forward
	toX, toY := fromX, fromY + direction
	if !chessboard.IsWithinBounds(toX, toY) {
		return nil, nil // Shouldn't error if out of bounds
	}

	isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)

	if err != nil {
		return nil, errors.New("failed to check if square is empty")
	}

	if isSquareEmpty {
		if toY == promotionRank {
			for _, promotionPiece := range []int{2, 3, 4, 5} { // Create a move for each piece it can promote too
				moves = append(moves, entity.NewMoveEntity(
					HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
					HelperService.IntPtr(toX), HelperService.IntPtr(toY),
					HelperService.IntPtr(promotionPiece),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				))
			}
		} else {
			moves = append(moves, entity.NewMoveEntity(
				HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
				HelperService.IntPtr(toX), HelperService.IntPtr(toY),
				HelperService.IntPtr(0),
				HelperService.BoolPtr(false), HelperService.BoolPtr(false),
				HelperService.IntPtr(0),
			))
		}
	}

	// 2 moves forward

	if fromY == startRank {
		toX, toY := fromX, fromY + (direction * 2)
		if !chessboard.IsWithinBounds(toX, toY) {
			return nil, nil // Shouldn't error if out of bounds
		}

		isSquareEmpty, err := chessboard.IsSquareEmpty(toY, toX)

		if err != nil {
			return nil, errors.New("failed to check if square is empty")
		}

		if isSquareEmpty { // ... should always be inbounds if from start position, but we'll check just in case
			moves = append(moves, entity.NewMoveEntity(
				HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
				HelperService.IntPtr(toX), HelperService.IntPtr(toY),
				HelperService.IntPtr(0),
				HelperService.BoolPtr(false), HelperService.BoolPtr(false),
				HelperService.IntPtr(0),
			))
		}
	}

	// Capture diagonal left

	for _, deltaX := range []int{-1, 1} {
		toX, toY = fromX + deltaX, fromY + direction
		if !chessboard.IsWithinBounds(toX, toY) {
			break // Shouldn't error if out of bounds
		}
		
		IsOpponent, err := chessboard.IsOpponent(piece, toY, toX)

		if err != nil {
			return nil, errors.New("failed to check if square is an opponent")
		}

		if IsOpponent {
			capturedPiece, err := chessboard.GetPiece(toY, toX) 

			if err != nil {
				return nil, errors.New("failed to get captured piece")
			}

			moves = append(moves, entity.NewMoveEntity(
				HelperService.IntPtr(fromX), HelperService.IntPtr(fromY),
				HelperService.IntPtr(toX), HelperService.IntPtr(toY),
				HelperService.IntPtr(0),
				HelperService.BoolPtr(false), HelperService.BoolPtr(false),
				HelperService.IntPtr(capturedPiece),
			))
		}
	}

	return moves, nil
}

// func getKnightMove (piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
// 	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
// 	move := entity.NewMoveEntity(
// 		&fromX, 
// 		&fromY, 
// 		&toX, 
// 		&toY, 
// 		&promotion, 
// 		&isCastling,
// 		&isEnPassant, 
// 		&captured,
// 	)
// 	var moves []entity.MoveEntityInterface
// 	moves = append(moves, move)
// 	return moves, nil
// }

// func getBishopMove (piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
// 	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
// 	move := entity.NewMoveEntity(
// 		&fromX, 
// 		&fromY, 
// 		&toX, 
// 		&toY, 
// 		&promotion, 
// 		&isCastling,
// 		&isEnPassant, 
// 		&captured,
// 	)
// 	var moves []entity.MoveEntityInterface
// 	moves = append(moves, move)
// 	return moves, nil
// }

// func getRookMove (piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
// 	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
// 	move := entity.NewMoveEntity(
// 		&fromX, 
// 		&fromY, 
// 		&toX, 
// 		&toY, 
// 		&promotion, 
// 		&isCastling,
// 		&isEnPassant, 
// 		&captured,
// 	)
// 	var moves []entity.MoveEntityInterface
// 	moves = append(moves, move)
// 	return moves, nil
// }

// func getQueenMove (piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
// 	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
// 	move := entity.NewMoveEntity(
// 		&fromX, 
// 		&fromY, 
// 		&toX, 
// 		&toY, 
// 		&promotion, 
// 		&isCastling,
// 		&isEnPassant, 
// 		&captured,
// 	)
// 	var moves []entity.MoveEntityInterface
// 	moves = append(moves, move)
// 	return moves, nil
// }

// func getKingMove (piece int, fromX int, fromY int, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
// 	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
// 	move := entity.NewMoveEntity(
// 		&fromX, 
// 		&fromY, 
// 		&toX, 
// 		&toY, 
// 		&promotion, 
// 		&isCastling,
// 		&isEnPassant, 
// 		&captured,
// 	)
// 	var moves []entity.MoveEntityInterface
// 	moves = append(moves, move)
// 	return moves, nil
// }
