package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/gabrielg2020/chess-api/api/entity"
)

type MockFENService struct {
	mock.Mock
}

func (m *MockFENService) Validate(fen string) (error) {
	args := m.Called(fen)
	return args.Error(0)
}

func (m *MockFENService) Parse(validFen string) (entity.ChessboardEntityInterface, error) {
	args := m.Called(validFen)
	return args.Get(0).(entity.ChessboardEntityInterface), args.Error(1)
}