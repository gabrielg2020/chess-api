package MoveService

import (
	"errors"
	"math"

	"github.com/gabrielg2020/chess-api/api/entity"
)

type MoveServiceInterface interface {
	FindBestMove(chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error)
}

type MoveService struct{}

func NewMoveService() *MoveService {
	return &MoveService{}
}

func (service *MoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	// 1. Find Psudo Legal Moves
		// a. Create a moves array
	var moves []entity.MoveEntityInterface
	board, err := chessboard.GetBoard()
	if err == nil {
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
				getPawnMove(piece, row, col)
			case 2: // Get Knight Move
				getKnightMove(piece, row, col)
			case 3: // Get Bishop Move
				getBishopMove(piece, row, col)
			case 4: // Get Rook Move
				getRookMove(piece, row, col)
			case 5: // Get Queen Move
				getQueenMove(piece, row, col)
			case 6: // Get King Move
				getKingMove(piece, row, col)
			}
		}
	}
		// c. For each piece, find all moves
	// 2. Filter for Legal Moves
		// a. Remove any move that goes off the board
		// b. Remove any move that place king in check
	// 3. Return Legal Moves
		// a. Return moves array
	return moves, nil
}
func getPawnMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {

	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func getKnightMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {
	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func getBishopMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {
	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func getRookMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {
	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func getQueenMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {
	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func getKingMove (piece int, fromX int, fromY int) ([]entity.MoveEntityInterface, error) {
	toX, toY, promotion, isCastling, isEnPassant, captured := 1, 1, 1, true, true, 1
	move := entity.NewMoveEntity(
		&fromX, 
		&fromY, 
		&toX, 
		&toY, 
		&promotion, 
		&isCastling,
		&isEnPassant, 
		&captured,
	)
	var moves []entity.MoveEntityInterface
	moves = append(moves, move)
	return moves, nil
}

func isMoveOutOfBounds (move entity.MoveEntityInterface) (error) {
	toX, errX := move.GetToX()
	toY, errY := move.GetToY()

	if (errX != nil) || (errY != nil) {
		return errors.New("failed to retrive toX or toY")
	}

	if (toX > 7) || (toX < 0) || (toY > 7) || (toY < 0) {
		return errors.New("move out of bounds")
	} else {
		return nil
	}
}