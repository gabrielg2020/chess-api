package entity

import (
	"testing"
	"errors"
	
	"github.com/stretchr/testify/assert"
)
func Test_ChessboardEntity_GetBoard(t *testing.T) {
	testCases := []struct {
		name              string
		board             [8][8]int
		expectedResponse  [8][8]int
		expectedError     error
	}{
		{
			name: "chessboard.board is set",
			board: [8][8]int{
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
			board: [8][8]int{
				{-7, -7, -7, -7, -7, -7, -7, -7},
				{-7, -7, -7, -7, -7, -7, -7, -7},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{7, 7, 7, 7, 7, 7, 7, 7},
				{7, 7, 7, 7, 7, 7, 7, 7},
			},
			expectedResponse: [8][8]int{},
			expectedError: errors.New("chessboard.board is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(tc.board, "fen", "", "", "", "", "")
			// Act
			response, err := entity.GetBoard()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_ChessboardEntity_GetFen(t *testing.T) {
	testCases := []struct {
		name              string
		fen               string
		expectedResponse  string
		expectedError     error
	}{
		{
			name: "chessboard.fen is set",
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedResponse: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedError: nil,
		},
		{
			name: "chessboard.fen is not set",
			fen: "",
			expectedResponse: "",
			expectedError: errors.New("chessboard.fen is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity([8][8]int{}, tc.fen, "", "", "", "", "")
			// Act
			response, err := entity.GetFen()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_ChessboardEntity_GetActiveColour(t *testing.T) {
	testCases := []struct {
		name              string
		activeColour      string
		expectedResponse  string
		expectedError     error
	}{
		{
			name: "chessboard.activeColour is set",
			activeColour: "w",
			expectedResponse: "w",
			expectedError: nil,
		},
		{
			name: "chessboard.activeColour is not set",
			activeColour: "",
			expectedResponse: "",
			expectedError: errors.New("chessboard.activeColour is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity([8][8]int{}, "", tc.activeColour, "", "", "", "")
			// Act
			response, err := entity.GetActiveColour()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_ChessboardEntity_GetCastlingRights(t *testing.T) {
	testCases := []struct {
		name              string
		castlingRights      string
		expectedResponse  string
		expectedError     error
	}{
		{
			name: "chessboard.castlingRights is set",
			castlingRights: "w",
			expectedResponse: "w",
			expectedError: nil,
		},
		{
			name: "chessboard.castlingRights is not set",
			castlingRights: "",
			expectedResponse: "",
			expectedError: errors.New("chessboard.castlingRights is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity([8][8]int{}, "", "", tc.castlingRights, "", "", "")
			// Act
			response, err := entity.GetCastlingRights()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_ChessboardEntity_GetEnPassantSquare(t *testing.T) {
	testCases := []struct {
		name              string
		enPassantSquare      string
		expectedResponse  string
		expectedError     error
	}{
		{
			name: "chessboard.enPassantSquare is set",
			enPassantSquare: "w",
			expectedResponse: "w",
			expectedError: nil,
		},
		{
			name: "chessboard.enPassantSquare is not set",
			enPassantSquare: "",
			expectedResponse: "",
			expectedError: errors.New("chessboard.enPassantSquare is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity([8][8]int{}, "", "", "", tc.enPassantSquare, "", "")
			// Act
			response, err := entity.GetEnPassantSquare()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}