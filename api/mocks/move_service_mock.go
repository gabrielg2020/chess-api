package mocks

import (
	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/stretchr/testify/mock"
)

type MockMoveService struct {
	mock.Mock
}

func (m *MockMoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error) {
	args := m.Called(chessboard)
	return args.String(0), args.Error(1)
}
