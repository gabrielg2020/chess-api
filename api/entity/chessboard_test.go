package entity

import (
	"testing"
	"errors"
	
	"github.com/stretchr/testify/assert"
)
func Test_ChessboardEntity_GetBoard(t *testing.T) {
	testCases := []struct {
		name             string
		board            *[8][8]int
		expectedResponse [8][8]int
		expectedError    error
	}{
		{
			name: "chessboard.board is set",
			board: &[8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			},
			expectedResponse: [8][8]int{
				{-4, -2, -3, -5, -6, -3, -2, -4},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{4, 2, 3, 5, 6, 3, 2, 4},
			},
			expectedError: nil,
		},
		{
			name: "chessboard.board is not set",
			board: nil,
			expectedResponse: [8][8]int{},
			expectedError: errors.New("chessboard.board is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(tc.board, nil, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetBoard()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetFen(t *testing.T) {
	testCases := []struct {
		name             string
		fen              *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.fen is set",
			fen: func() *string {
				s := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
				return &s
			}(),
			expectedResponse: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: nil,
		},
		{
			name: "chessboard.fen is not set",
			fen: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.fen is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, tc.fen, nil, nil, nil, nil, nil)
			// Act
			response, err := entity.GetFen()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetActiveColour(t *testing.T) {
	testCases := []struct {
		name             string
		activeColour     *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.activeColour is set",
			activeColour: func() *string {
				s := "w"
				return &s
			}(),
			expectedResponse: "w",
			expectedError: nil,
		},
		{
			name: "chessboard.activeColour is not set",
			activeColour: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.activeColour is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, tc.activeColour, nil, nil, nil, nil)
			// Act
			response, err := entity.GetActiveColour()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetCastlingRights(t *testing.T) {
	testCases := []struct {
		name             string
		castlingRights   *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.castlingRights is set",
			castlingRights: func() *string {
				s := "KQkq"
				return &s
			}(),
			expectedResponse: "KQkq",
			expectedError: nil,
		},
		{
			name: "chessboard.castlingRights is not set",
			castlingRights: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.castlingRights is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, tc.castlingRights, nil, nil, nil)
			// Act
			response, err := entity.GetCastlingRights()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetEnPassantSquare(t *testing.T) {
	testCases := []struct {
		name             string
		enPassantSquare  *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.enPassantSquare is set",
			enPassantSquare: func() *string {
				s := "e3"
				return &s
			}(),
			expectedResponse: "e3",
			expectedError: nil,
		},
		{
			name: "chessboard.enPassantSquare is not set",
			enPassantSquare: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.enPassantSquare is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, tc.enPassantSquare, nil, nil)
			// Act
			response, err := entity.GetEnPassantSquare()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetHalfmoveClock(t *testing.T) {
	testCases := []struct {
		name             string
		halfmoveClock    *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.halfmoveClock is set",
			halfmoveClock: func() *string {
				s := "0"
				return &s
			}(),
			expectedResponse: "0",
			expectedError: nil,
		},
		{
			name: "chessboard.halfmoveClock is not set",
			halfmoveClock: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.halfmoveClock is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, tc.halfmoveClock, nil)
			// Act
			response, err := entity.GetHalfmoveClock()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ChessboardEntity_GetFullmoveNumber(t *testing.T) {
	testCases := []struct {
		name             string
		fullmoveNumber   *string
		expectedResponse string
		expectedError    error
	}{
		{
			name: "chessboard.fullmoveNumber is set",
			fullmoveNumber: func() *string {
				s := "1"
				return &s
			}(),
			expectedResponse: "1",
			expectedError:    nil,
		},
		{
			name: "chessboard.fullmoveNumber is not set",
			fullmoveNumber: nil,
			expectedResponse: "",
			expectedError: errors.New("chessboard.fullmoveNumber is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(nil, nil, nil, nil, nil, nil, tc.fullmoveNumber)
			// Act
			response, err := entity.GetFullmoveNumber()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}