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
			name:  "Bishop moves to empty squares (Blocked by own piece)",
			piece: 3,
			fromY: 4, fromX: 3,
			deltaXs:   []int{1, -1, 1, -1},
			deltaYs:   []int{1, 1, -1, -1},
			isSliding: true,
			setupMock: func(m *mocks.MockChessboardEntity) {
				positions := []struct{ toY, toX int }{
					{5, 4}, {6, 5}, {5, 2}, {6, 1},
					{3, 4}, {2, 5}, {3, 2}, {2, 1},
				}
				for i := 1; i < len(positions); i += 2 {
					pos1 := positions[i-1]
					pos2 := positions[i]
					toY1, toX1 := pos1.toY, pos1.toX
					m.On("IsWithinBounds", toY1, toX1).Return(true)
					m.On("IsSquareEmpty", toY1, toX1).Return(true, nil)
					toY2, toX2 := pos2.toY, pos2.toX
					m.On("IsWithinBounds", toY2, toX2).Return(true)
					m.On("IsSquareEmpty", toY2, toX2).Return(false, nil)
					m.On("IsOpponent", 3, toY2, toX2).Return(true, nil)
					m.On("GetPiece", toY2, toX2).Return(-1, nil)
				}
			},
			expectedMoves: []entity.MoveEntityInterface{
				newMockMoveEntity(3, 4, 4, 5, 0, false, false, 0),
				newMockMoveEntity(3, 4, 5, 6, 0, false, false, -1),
				newMockMoveEntity(3, 4, 2, 5, 0, false, false, 0),
				newMockMoveEntity(3, 4, 1, 6, 0, false, false, -1),
				newMockMoveEntity(3, 4, 4, 3, 0, false, false, 0),
				newMockMoveEntity(3, 4, 5, 2, 0, false, false, -1),
				newMockMoveEntity(3, 4, 2, 3, 0, false, false, 0),
				newMockMoveEntity(3, 4, 1, 2, 0, false, false, -1),
			},
			expectedError: nil,
		},
		{
			name:  "Fail to add move when is NOT sliding",
			piece: 2,
			fromY: 4, fromX: 4,
			deltaXs:   []int{1, 1, -1, -1, 2, 2, -2, -2},
			deltaYs:   []int{2, -2, 2, -2, 1, -1, 1, -1},
			isSliding: false,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsWithinBounds", 6, 5).Return(true)
				m.On("IsSquareEmpty", 6, 5).Return(false, errors.New("test error in IsSquareEmpty"))
			},
			expectedMoves: []entity.MoveEntityInterface{},
			expectedError: errors.New("MoveService.generateMoves: MoveService.tryAddMove: test error in IsSquareEmpty"),
		},
		{
			name:  "Fail to add move when is sliding",
			piece: 3,
			fromY: 4, fromX: 3,
			deltaXs:   []int{1, -1, 1, -1},
			deltaYs:   []int{1, 1, -1, -1},
			isSliding: true,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsWithinBounds", 5, 4).Return(true)
				m.On("IsSquareEmpty", 5, 4).Return(false, errors.New("test error in IsSquareEmpty"))
			},
			expectedMoves: []entity.MoveEntityInterface{},
			expectedError: errors.New("MoveService.generateMoves: MoveService.tryAddMove: test error in IsSquareEmpty"),
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
			expectedResponse: false,
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

func Test_MoveService_canCastleKingSide(t *testing.T) {
	testCases := []struct {
		name             string
		fromX, fromY     int
		setupMock        func(m *mocks.MockChessboardEntity)
		expectedResponse bool
		expectedError    error
	}{
		{
			name:  "Can castle king side",
			fromX: 4,
			fromY: 0,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 0, 5).Return(true, nil)
				m.On("IsSquareEmpty", 0, 6).Return(true, nil)
				m.On("GetPiece", 0, 7).Return(4, nil)
			},
			expectedResponse: true,
			expectedError:    nil,
		},
		{
			name:  "Cannot castle king side (Space between king and rook is not empty)",
			fromX: 4,
			fromY: 0,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 0, 5).Return(false, nil)
			},
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:  "Cannot castle king side (Rook is not in the correct position)",
			fromX: 4,
			fromY: 0,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 0, 5).Return(true, nil)
				m.On("IsSquareEmpty", 0, 6).Return(true, nil)
				m.On("GetPiece", 0, 7).Return(0, nil)
			},
			expectedResponse: false,
			expectedError:    nil,
		},
		{
			name:  "Error in IsSquareEmpty",
			fromX: 4,
			fromY: 0,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 0, 5).Return(false, errors.New("test error in IsSquareEmpty"))
			},
			expectedResponse: false,
			expectedError:    errors.New("MoveService.canCastleKingSide: test error in IsSquareEmpty"),
		},
		{
			name:  "Error in GetPiece",
			fromX: 4,
			fromY: 0,
			setupMock: func(m *mocks.MockChessboardEntity) {
				m.On("IsSquareEmpty", 0, 5).Return(true, nil)
				m.On("IsSquareEmpty", 0, 6).Return(true, nil)
				m.On("GetPiece", 0, 7).Return(0, errors.New("test error in GetPiece"))
			},
			expectedResponse: false,
			expectedError:    errors.New("MoveService.canCastleKingSide: test error in GetPiece"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mocks.MockChessboardEntity)
			tc.setupMock(mockChessboard)

			// Act
			result, err := canCastleKingSide(6, tc.fromY, tc.fromX, mockChessboard)

			// Assert
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, result)
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
