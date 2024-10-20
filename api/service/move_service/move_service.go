package MoveService

import (
	"errors"

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