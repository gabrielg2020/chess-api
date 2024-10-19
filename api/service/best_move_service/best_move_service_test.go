package BestMoveService

import (
	"testing"
	// "errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Entity
type mockChessboardEntity struct {
	mock.Mock
}

func (m *mockChessboardEntity) GetFen() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockChessboardEntity) GetBoard() ([8][8]int, error) {
	args := m.Called()
	return args.Get(0).([8][8]int), args.Error(1)
}

func Test_BestMoveService_FindBestMoveWithArray(t *testing.T) {
	service := NewBestMoveService()

	testCases := []struct {
		name          string
		setupMock     func(m *mockChessboardEntity)
		expectedMove  string
		expectedError error
	}{
		{
			// TODO Add test cases when functionallity is made
			name: "Test Case 1",
			setupMock: func(m *mockChessboardEntity) {
				m.On("GetBoard").Return([8][8]int{
					{4, 2, 3, 5, 6, 3, 2, 4},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1},
					{-4, -2, -3, -5, -6, -3, -2, -4},
				}, nil)
			},
			expectedMove: "a2a4",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mockChessboardEntity)
			tc.setupMock(mockChessboard)

			// Act
			move, err := service.FindBestMove(mockChessboard)

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedMove, move)
		})
	}
}
