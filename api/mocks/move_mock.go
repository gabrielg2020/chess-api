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

func (m *MockMoveEntity) isCastling() (bool, error) {
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