package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockChessboardEntity struct {
	mock.Mock
}

func (m *MockChessboardEntity) GetBoard() ([8][8]int, error) {
	args := m.Called()
	return args.Get(0).([8][8]int), args.Error(1)
}

func (m *MockChessboardEntity) GetFen() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetActiveColour() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetCastlingRights() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetEnPassantSquare() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetHalfmoveClock() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetFullmoveNumber() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockChessboardEntity) GetPiece(row int, col int) (int, error) {
	args := m.Called(row, col)
	return args.Int(0), args.Error(1)
}

func (m *MockChessboardEntity) IsSquareEmpty(row int, col int) (bool, error) {
	args := m.Called(row, col)
	return args.Bool(0), args.Error(1)
}

func (m *MockChessboardEntity) IsOpponent(piece int, row int, col int) (bool, error) {
	args := m.Called(piece, row, col)
	return args.Bool(0), args.Error(1)
}

func (m *MockChessboardEntity) IsWithinBounds(row int, col int) bool {
	args := m.Called(row, col)
	return args.Bool(0)
}

func (m *MockChessboardEntity) IsSquareUnderAttack(piece int, row int, col int) (bool, error) {
	args := m.Called(piece, row, col)
	return args.Bool(0), args.Error(1)
}
