package mocks

import (
	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/stretchr/testify/mock"
)

type MockMoveService struct {
	mock.Mock
}

func (m *MockMoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (entity.MoveEntityInterface, error) {
	args := m.Called(chessboard)
	return args.Get(0).(entity.MoveEntityInterface), args.Error(1)
}

func (m *MockMoveService) FindPseudoLegalMoves(colour string, chessboard entity.ChessboardEntityInterface) ([]entity.MoveEntityInterface, error) {
	args := m.Called(colour, chessboard)
	return args.Get(0).([]entity.MoveEntityInterface), args.Error(1)
}
