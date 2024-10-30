package MoveService

import (
	"errors"
	"github.com/gabrielg2020/chess-api/api/entity"
	"github.com/gabrielg2020/chess-api/api/mocks"
	"github.com/gabrielg2020/chess-api/api/service/helper_service"
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

func compareMoves(t *testing.T, expected entity.MoveEntityInterface, actual entity.MoveEntityInterface) {
	// Compare fromX
	fromXExpected, err := expected.GetFromX()
	assert.NoError(t, err)
	fromXActual, err := actual.GetFromX()
	assert.NoError(t, err)
	assert.Equal(t, fromXExpected, fromXActual)
	// Compare fromY
	fromYExpected, err := expected.GetFromY()
	assert.NoError(t, err)
	fromYActual, err := actual.GetFromY()
	assert.NoError(t, err)
	assert.Equal(t, fromYExpected, fromYActual)
	// Compare toX
	toXExpected, err := expected.GetToX()
	assert.NoError(t, err)
	toXActual, err := actual.GetToX()
	assert.NoError(t, err)
	assert.Equal(t, toXExpected, toXActual)
	// Compare toY
	toYExpected, err := expected.GetToY()
	assert.NoError(t, err)
	toYActual, err := actual.GetToY()
	assert.NoError(t, err)
	assert.Equal(t, toYExpected, toYActual)
	// Compare promotion
	promotionExpected, err := expected.GetPromotion()
	assert.NoError(t, err)
	promotionActual, err := actual.GetPromotion()
	assert.NoError(t, err)
	assert.Equal(t, promotionExpected, promotionActual)
	// Compare isCastling
	isCastlingExpected, err := expected.IsCastling()
	assert.NoError(t, err)
	isCastlingActual, err := actual.IsCastling()
	assert.NoError(t, err)
	assert.Equal(t, isCastlingExpected, isCastlingActual)
	// Compare isEnPassant
	isEnPassantExpected, err := expected.IsEnPassant()
	assert.NoError(t, err)
	isEnPassantActual, err := actual.IsEnPassant()
	assert.NoError(t, err)
	assert.Equal(t, isEnPassantExpected, isEnPassantActual)
	// Compare getCaptured
	capturedExpected, err := expected.GetCaptured()
	assert.NoError(t, err)
	capturedActual, err := actual.GetCaptured()
	assert.NoError(t, err)
	assert.Equal(t, capturedExpected, capturedActual)
}

func Test_MoveService_getPawnMove(t *testing.T) {
	testCases := []struct {
		name          string
		piece         int
		fromX         int // col
		fromY         int // row
		setupMock     func(m *mocks.MockChessboardEntity)
		expectedMoves []entity.MoveEntityInterface
		expectedError error
	}{
		{
			name:  "White Pawn 1 Move Forward",
			piece: 1,
			fromX: 2,
			fromY: 3,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 2, 2).Return(true, nil)
				// Square 2 ahead is not empty
				m.On("IsSquareEmpty", 1, 2).Return(false, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", 1, 2, 1).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", 1, 2, 3).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(2), HelperService.IntPtr(2),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Black Pawn 1 Move Forward",
			piece: -1,
			fromX: 2,
			fromY: 3,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 4, 2).Return(true, nil)
				// Square 2 ahead is not empty
				m.On("IsSquareEmpty", 5, 2).Return(false, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", -1, 4, 3).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", -1, 4, 1).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(2), HelperService.IntPtr(4),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "White Pawn 2 Move Forward",
			piece: 1,
			fromX: 2,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 5, 2).Return(true, nil)
				// Square 2 ahead is empty
				m.On("IsSquareEmpty", 4, 2).Return(true, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", 1, 5, 3).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", 1, 5, 1).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(5),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(4),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(true),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Black Pawn 2 Move Forward",
			piece: -1,
			fromX: 2,
			fromY: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 2, 2).Return(true, nil)
				// Square 2 ahead is empty
				m.On("IsSquareEmpty", 3, 2).Return(true, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", -1, 2, 3).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", -1, 2, 1).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(2),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(true),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "White Pawn Promotion",
			piece: 1,
			fromX: 2,
			fromY: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 0, 2).Return(true, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", 1, 0, 1).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", 1, 0, 3).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(0),
					HelperService.IntPtr(2),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(0),
					HelperService.IntPtr(3),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(0),
					HelperService.IntPtr(4),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(2), HelperService.IntPtr(0),
					HelperService.IntPtr(5),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Black Pawn Promotion",
			piece: -1,
			fromX: 2,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is empty
				m.On("IsSquareEmpty", 7, 2).Return(true, nil)
				// Left, 1 ahead has no opponent
				m.On("IsOpponent", -1, 7, 1).Return(false, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", -1, 7, 3).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(7),
					HelperService.IntPtr(-2),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(7),
					HelperService.IntPtr(-3),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(7),
					HelperService.IntPtr(-4),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(2), HelperService.IntPtr(7),
					HelperService.IntPtr(-5),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(0),
				),
			},
			expectedError: nil,
		},
		{
			name:  "White Pawn Take Left and Right",
			piece: 1,
			fromX: 2,
			fromY: 3,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is not empty
				m.On("IsSquareEmpty", 2, 2).Return(false, nil)
				// Left, 1 ahead has opponent
				m.On("IsOpponent", 1, 2, 1).Return(true, nil)
				m.On("GetPiece", 2, 1).Return(-4, nil)
				// Right, 1 ahead has opponent
				m.On("IsOpponent", 1, 2, 3).Return(true, nil)
				m.On("GetPiece", 2, 3).Return(-4, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(1), HelperService.IntPtr(2),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(3), HelperService.IntPtr(2),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Black Pawn Take Left and Right",
			piece: -1,
			fromX: 2,
			fromY: 3,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is not empty
				m.On("IsSquareEmpty", 4, 2).Return(false, nil)
				// Left, 1 ahead has opponent
				m.On("IsOpponent", -1, 4, 1).Return(true, nil)
				m.On("GetPiece", 4, 1).Return(4, nil)
				// Right, 1 ahead has opponent
				m.On("IsOpponent", -1, 4, 3).Return(true, nil)
				m.On("GetPiece", 4, 3).Return(4, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(1), HelperService.IntPtr(4),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(3),
					HelperService.IntPtr(3), HelperService.IntPtr(4),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
			},
			expectedError: nil,
		},
		{
			name:  "White Pawn Take Left With Promotion",
			piece: 1,
			fromX: 2,
			fromY: 1,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is not empty
				m.On("IsSquareEmpty", 0, 2).Return(false, nil)
				// Left, 1 ahead has opponent
				m.On("IsOpponent", 1, 0, 1).Return(true, nil)
				m.On("GetPiece", 0, 1).Return(-4, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", 1, 0, 3).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(1), HelperService.IntPtr(0),
					HelperService.IntPtr(2),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(1), HelperService.IntPtr(0),
					HelperService.IntPtr(3),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(1), HelperService.IntPtr(0),
					HelperService.IntPtr(4),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(1),
					HelperService.IntPtr(1), HelperService.IntPtr(0),
					HelperService.IntPtr(5),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-4),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Black Pawn Take Left With Promotion",
			piece: -1,
			fromX: 2,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Square 1 ahead is not empty
				m.On("IsSquareEmpty", 7, 2).Return(false, nil)
				// Left, 1 ahead has opponent
				m.On("IsOpponent", -1, 7, 1).Return(true, nil)
				m.On("GetPiece", 7, 1).Return(4, nil)
				// Right, 1 ahead has no opponent
				m.On("IsOpponent", -1, 7, 3).Return(false, nil)
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(1), HelperService.IntPtr(7),
					HelperService.IntPtr(-2),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(1), HelperService.IntPtr(7),
					HelperService.IntPtr(-3),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(1), HelperService.IntPtr(7),
					HelperService.IntPtr(-4),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
				entity.NewMoveEntity(
					HelperService.IntPtr(2), HelperService.IntPtr(6),
					HelperService.IntPtr(1), HelperService.IntPtr(7),
					HelperService.IntPtr(-5),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(4),
				),
			},
			expectedError: nil,
		},
		{
			name:  "Failed To Check If Square Is Empty When Checking 1 Move Ahead",
			piece: 1,
			fromX: 3,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Fail on 1st square
				m.On("IsSquareEmpty", 5, 3).Return(false, errors.New("chessboard.board is not set"))
			},
			expectedMoves: nil,
			expectedError: errors.New("failed to check if square is empty"),
		},
		{
			name:  "Failed To Check If Square Is Empty When Checking 2 Move Ahead",
			piece: 1,
			fromX: 3,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Pass on 1st square
				m.On("IsSquareEmpty", 5, 3).Return(true, nil)
				// Fail on 2nd square
				m.On("IsSquareEmpty", 4, 3).Return(false, errors.New("chessboard.board is not set"))
			},
			expectedMoves: nil,
			expectedError: errors.New("failed to check if square is empty"),
		},
		{
			name:  "Failed To Check If Is Opponent",
			piece: 1,
			fromX: 3,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Pass on 1st square
				m.On("IsSquareEmpty", 5, 3).Return(true, nil)
				// Pass on 2nd square
				m.On("IsSquareEmpty", 4, 3).Return(true, nil)
				// Fail Checking IsOpponent
				m.On("IsOpponent", 1, 5, 2).Return(false, errors.New("chessboard.board is not set"))
			},
			expectedMoves: nil,
			expectedError: errors.New("failed to check if is opponent"),
		},
		{
			name:  "Failed To Get Captured Piece",
			piece: 1,
			fromX: 3,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Pass on 1st square
				m.On("IsSquareEmpty", 5, 3).Return(true, nil)
				// Pass on 2nd square
				m.On("IsSquareEmpty", 4, 3).Return(true, nil)
				// Pass Checking IsOpponent
				m.On("IsOpponent", 1, 5, 2).Return(true, nil)
				// Failed Getting Piece
				m.On("GetPiece", 5, 2).Return(-7, errors.New("row or col out of bounds"))
			},
			expectedMoves: nil,
			expectedError: errors.New("failed to get captured piece"),
		},
		{
			name:  "En Passant Capture",
			piece: 1,
			fromX: 3,
			fromY: 6,
			setupMock: func(m *mocks.MockChessboardEntity) {
				// Pass on 1st square
				m.On("IsSquareEmpty", 5, 3).Return(true, nil)
				// Pass on 2nd square
				m.On("IsSquareEmpty", 4, 3).Return(true, nil)
				// Pass Checking IsOpponent
				m.On("IsOpponent", 1, 5, 2).Return(true, nil)
				// Pass Getting Piece
				m.On("GetPiece", 5, 2).Return(0, nil)
				// Fail Checking IsOpponent
				m.On("IsOpponent", 1, 5, 4).Return(false, errors.New("chessboard.board is not set"))
			},
			expectedMoves: []entity.MoveEntityInterface{
				entity.NewMoveEntity(
					HelperService.IntPtr(3), HelperService.IntPtr(6),
					HelperService.IntPtr(5), HelperService.IntPtr(2),
					HelperService.IntPtr(0),
					HelperService.BoolPtr(false), HelperService.BoolPtr(false),
					HelperService.IntPtr(-1),
				),
			},
			expectedError: errors.New("failed to check if is opponent"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockChessboard := new(mocks.MockChessboardEntity)
			tc.setupMock(mockChessboard)

			// Act
			moves, err := getPawnMove(tc.piece, tc.fromY, tc.fromX, mockChessboard)

			// Assert
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Len(t, moves, len(tc.expectedMoves))

				for index, move := range moves {
					expectedMove := tc.expectedMoves[index]
					compareMoves(t, expectedMove, move)
				}
			}
		})
	}
}
