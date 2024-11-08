package MoveService

import (
	"errors"
	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/gabrielg2020/chess-api/api/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
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

func Test_MoveService_generateMoves(t *testing.T) {
	testCases := []struct {
		name             string
		piece            int
		fromY, fromX     int
		deltaXs, deltaYs []int
		isSliding        bool
		setupMock        func(m *mocks.MockChessboardEntity)
		expectedMoves    []entity.MoveEntityInterface
		expectedError    error
	}{
		{
			name:  "Knight moves to empty squares",
			piece: 2,
			fromY: 4, fromX: 4,
			deltaXs:   []int{1, 1, -1, -1, 2, 2, -2, -2},
			deltaYs:   []int{2, -2, 2, -2, 1, -1, 1, -1},
			isSliding: false,
			setupMock: func(m *mocks.MockChessboardEntity) {
				positions := []struct{ toY, toX int }{
					{6, 5}, {2, 5}, {6, 3}, {2, 3}, {5, 6}, {3, 6}, {5, 2}, {3, 2},
				}
				for _, pos := range positions {
					toY, toX := pos.toY, pos.toX
					m.On("IsWithinBounds", toY, toX).Return(true)
					m.On("IsSquareEmpty", toY, toX).Return(true, nil)
				}
			},
			expectedMoves: []entity.MoveEntityInterface{
				newMockMoveEntity(4, 4, 5, 6, 0, false, false, 0),
				newMockMoveEntity(4, 4, 5, 2, 0, false, false, 0),
				newMockMoveEntity(4, 4, 3, 6, 0, false, false, 0),
				newMockMoveEntity(4, 4, 3, 2, 0, false, false, 0),
				newMockMoveEntity(4, 4, 6, 5, 0, false, false, 0),
				newMockMoveEntity(4, 4, 6, 3, 0, false, false, 0),
				newMockMoveEntity(4, 4, 2, 5, 0, false, false, 0),
				newMockMoveEntity(4, 4, 2, 3, 0, false, false, 0),
			},
			expectedError: nil,
		},
		{
			name:  "Bishop movement blocked by own piece",
			piece: 3,
			fromY: 4, fromX: 4,
			deltaXs:   []int{1, -1, 1, -1},
			deltaYs:   []int{1, 1, -1, -1},
			isSliding: true,
			setupMock: func(m *mocks.MockChessboardEntity) {

			},
			expectedMoves: []entity.MoveEntityInterface{},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mocks.MockChessboardEntity)
			tc.setupMock(mockChessboard)
			var moves []entity.MoveEntityInterface

			// Act
			moves, err := generateMoves(tc.piece, tc.fromY, tc.fromX, tc.deltaXs, tc.deltaYs, tc.isSliding, mockChessboard)

			// Assert
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assertMovesEqual(t, tc.expectedMoves, moves)
			}
		})
	}
}

func Test_MoveService_tryAddMove(t *testing.T) {
	testCases := []struct {
		name             string
		piece            int
		fromY, fromX     int
		toY, toX         int
		setupMock        func(m *mocks.MockChessboardEntity)
		expectedMoves    []entity.MoveEntityInterface
		expectedResponse bool
		expectedError    error
	}{
		{
			name:  "Empty square",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(true, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				newMockMoveEntity(1, 1, 1, 2, 0, false, false, 0),
			},
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name:  "Square occupied by opponent",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(false, nil)
				m.On("IsOpponent", 1, 2, 1).Return(true, nil)
				m.On("GetPiece", 2, 1).Return(-3, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				newMockMoveEntity(1, 1, 1, 2, 0, false, false, -3),
			},
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name:  "Square occupied by own piece",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(false, nil)
				m.On("IsOpponent", 1, 2, 1).Return(false, nil)
			},
			expectedMoves:    []entity.MoveEntityInterface{},
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:  "Error in IsSquareEmpty",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(false, errors.New("test error in IsSquareEmpty"))
			},
			expectedMoves:    nil,
			expectedResponse: false,
			expectedError:    errors.New("MoveService.tryAddMove: test error in IsSquareEmpty"),
		},
		{
			name:  "Error in IsOpponent",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(false, nil)
				m.On("IsOpponent", 1, 2, 1).Return(false, errors.New("test error in IsOpponent"))
			},
			expectedMoves:    []entity.MoveEntityInterface{},
			expectedResponse: false,
			expectedError:    errors.New("MoveService.tryAddMove: test error in IsOpponent"),
		},
		{
			name:  "Square occupied by own piece",
			piece: 1,
			fromY: 1, fromX: 1,
			toY: 2, toX: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 2, 1).Return(false, nil)
				m.On("IsOpponent", 1, 2, 1).Return(true, nil)
				m.On("GetPiece", 2, 1).Return(-7, errors.New("test error in GetPiece"))
			},
			expectedMoves:    []entity.MoveEntityInterface{},
			expectedResponse: false,
			expectedError:    errors.New("MoveService.tryAddMove: test error in GetPiece"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mocks.MockChessboardEntity)
			tc.setupMock(mockChessboard)
			var moves []entity.MoveEntityInterface

			// Act
			result, err := tryAddMove(tc.piece, tc.fromY, tc.fromX, tc.toY, tc.toX, &moves, mockChessboard)

			// Assert
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, result)
				assertMovesEqual(t, tc.expectedMoves, moves)
			}
		})
	}
}

// Helper functions
func assertMovesEqual(t *testing.T, expected, actual []entity.MoveEntityInterface) {
	assert.Equal(t, len(expected), len(actual), "Number of moves should be equal")
	for i := range expected {
		assertMoveEqual(t, expected[i], actual[i])
	}
}

func assertMoveEqual(t *testing.T, expected, actual entity.MoveEntityInterface) {
	// Get FromX
	expectedFromX, err := expected.GetFromX()
	assert.NoError(t, err, "Expected move GetFromX should not return an error")
	actualFromX, err := actual.GetFromX()
	assert.NoError(t, err, "Actual move GetFromX should not return an error")
	assert.Equal(t, expectedFromX, actualFromX, "FromX should be equal")

	// Get FromY
	expectedFromY, err := expected.GetFromY()
	assert.NoError(t, err, "Expected move GetFromY should not return an error")
	actualFromY, err := actual.GetFromY()
	assert.NoError(t, err, "Actual move GetFromY should not return an error")
	assert.Equal(t, expectedFromY, actualFromY, "FromY should be equal")

	// Get ToX
	expectedToX, err := expected.GetToX()
	assert.NoError(t, err, "Expected move GetToX should not return an error")
	actualToX, err := actual.GetToX()
	assert.NoError(t, err, "Actual move GetToX should not return an error")
	assert.Equal(t, expectedToX, actualToX, "ToX should be equal")

	// Get ToY
	expectedToY, err := expected.GetToY()
	assert.NoError(t, err, "Expected move GetToY should not return an error")
	actualToY, err := actual.GetToY()
	assert.NoError(t, err, "Actual move GetToY should not return an error")
	assert.Equal(t, expectedToY, actualToY, "ToY should be equal")

	// Get Promotion
	expectedPromotion, err := expected.GetPromotion()
	assert.NoError(t, err, "Expected move GetPromotion should not return an error")
	actualPromotion, err := actual.GetPromotion()
	assert.NoError(t, err, "Actual move GetPromotion should not return an error")
	assert.Equal(t, expectedPromotion, actualPromotion, "Promotion should be equal")

	// IsCastling
	expectedIsCastling, err := expected.IsCastling()
	assert.NoError(t, err, "Expected move IsCastling should not return an error")
	actualIsCastling, err := actual.IsCastling()
	assert.NoError(t, err, "Actual move IsCastling should not return an error")
	assert.Equal(t, expectedIsCastling, actualIsCastling, "IsCastling should be equal")

	// IsEnPassant
	expectedIsEnPassant, err := expected.IsEnPassant()
	assert.NoError(t, err, "Expected move IsEnPassant should not return an error")
	actualIsEnPassant, err := actual.IsEnPassant()
	assert.NoError(t, err, "Actual move IsEnPassant should not return an error")
	assert.Equal(t, expectedIsEnPassant, actualIsEnPassant, "IsEnPassant should be equal")

	// GetCaptured
	expectedCaptured, err := expected.GetCaptured()
	assert.NoError(t, err, "Expected move GetCaptured should not return an error")
	actualCaptured, err := actual.GetCaptured()
	assert.NoError(t, err, "Actual move GetCaptured should not return an error")
	assert.Equal(t, expectedCaptured, actualCaptured, "Captured piece should be equal")
}

func newMockMoveEntity(fromX, fromY, toX, toY, promotion int, isCastling, isEnPassant bool, captured int) entity.MoveEntityInterface {
	mockMove := new(mocks.MockMoveEntity)

	mockMove.On("GetFromX").Return(fromX, nil)
	mockMove.On("GetFromY").Return(fromY, nil)
	mockMove.On("GetToX").Return(toX, nil)
	mockMove.On("GetToY").Return(toY, nil)
	mockMove.On("GetPromotion").Return(promotion, nil)
	mockMove.On("IsCastling").Return(isCastling, nil)
	mockMove.On("IsEnPassant").Return(isEnPassant, nil)
	mockMove.On("GetCaptured").Return(captured, nil)
	mockMove.On("GetChessNotation").Return("", nil)

	return mockMove
}
