package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/gabrielg2020/chess-api/api/entity"
)

type MockBestMoveService struct {
	mock.Mock
}

func (m *MockBestMoveService) FindBestMove(chessboard entity.ChessboardEntityInterface) (string, error) {
	args := m.Called(chessboard)
	return args.String(0), args.Error(1)
}