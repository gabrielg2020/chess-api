package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockMoveEntity struct {
	mock.Mock
}

func (m *MockMoveEntity) GetFromX() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) GetFromY() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) GetToX() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) GetToY() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) GetPromotion() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) IsCastling() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockMoveEntity) IsEnPassant() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockMoveEntity) GetCaptured() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockMoveEntity) GetChessNotation() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockMoveEntity) GetMoveProperties() (int, int, int, int, int, bool, bool, int, error) {
	args := m.Called()
	return args.Int(0), args.Int(1), args.Int(2), args.Int(3), args.Int(4), args.Bool(5), args.Bool(6), args.Int(7), args.Error(8)
}
