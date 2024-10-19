package BestMoveService

import "github.com/gabrielg2020/chess-api/api/entity"

type BestMoveServiceInterface interface {
	FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error)
}

type BestMoveService struct {}

func NewBestMoveService() *BestMoveService {
	return &BestMoveService{}
}

func (service *BestMoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error){
	// TODO Create best move functionallity
	return "a2a4", nil
}

