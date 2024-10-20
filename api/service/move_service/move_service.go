package MoveService

import "github.com/gabrielg2020/chess-api/api/entity"

type MoveServiceInterface interface {
	FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error)
}

type MoveService struct{}

func NewMoveService() *MoveService {
	return &MoveService{}
}

func (service *MoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error) {
	// TODO Create best move functionallity
	return "a2a4", nil
}
