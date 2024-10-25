package MoveService

import (
	"testing"
	// "errors"

	"github.com/gabrielg2020/chess-api/api/mocks"
	// "github.com/gabrielg2020/chess-api/api/entity"
	// "github.com/stretchr/testify/assert"
)

func Test_MoveService_FindBestMove(t *testing.T) {
	// TODO [FindBestMove] Complete test when move_service.FindBestMove is completed.
	// service := NewMoveService()

	testCases := []struct {
		name          string
		setupMock     func(m *mocks.MockChessboardEntity)
		expectedMove  string
		expectedError error
	}{
		{
			// TODO [FindBestMove] Complete test when move_service.FindBestMove is completed.
			name: "Test Case 1",
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("GetBoard").Return([8][8]int{
					{-4, -2, -3, -5, -6, -3, -2, -4},
					{-1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1},
					{4, 2, 3, 5, 6, 3, 2, 4},
				}, nil)
			},
			expectedMove:  "a2a4",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mocks.MockChessboardEntity)
			tc.setupMock(mockChessboard)

			// TODO [FindBestMove] Complete test when move_service.FindBestMove is completed.
			// Act
			//move, err := service.FindBestMove(mockChessboard)

			// Assert
			// assert.Equal(t, tc.expectedError, err)
			// assert.Equal(t, tc.expectedMove, move)
		})
	}
}

// func Test_MoveService_getPawnMove(t *testing.T){
// 	testCases := []struct {
// 		name                    string
// 		setupMock               func(m *mocks.MockChessboardEntity)
// 		expectedMoves           func(moves *[]mocks.MockMoveEntity)
// 		expectedError           error
// 	}{
// 		{
// 			name: "Get All Moves For A White Pawn",
// 			setupMock: func(m *mocks.MockChessboardEntity) {
// 				m.On("GetBoard").Return([8][8]int{
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 1, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 				}, nil)
// 			},
// 			expectedMoves: func(moves *[]mocks.MockMoveEntity) {

// 			}
// 		}
// 	}
// }
