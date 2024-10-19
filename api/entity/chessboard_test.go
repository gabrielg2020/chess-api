package entity

import (
	"testing"
	"errors"
	
	"github.com/stretchr/testify/assert"
)

func Test_ChessboardEntity_GetFen(t *testing.T) {
	testCases := []struct {
		name              string
		fen               string
		expectedResponse string
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
			entity := NewChessboardEntity([8][8]int{}, tc.fen, "", "", "", 0, 0)
			// Act
			response, err := entity.GetFen()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_ChessboardEntity_GetBoard(t *testing.T) {
	testCases := []struct {
		name              string
		board             [8][8]int
		expectedResponse [8][8]int
		expectedError     error
	}{
		{
			name: "chessboard.board is set",
			board: [8][8]int{
				{4, 2, 3, 5, 6, 3, 2, 4},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{-4, -2, -3, -5, -6, -3, -2, -4},
			},
			expectedResponse: [8][8]int{
				{4, 2, 3, 5, 6, 3, 2, 4},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{-1, -1, -1, -1, -1, -1, -1, -1},
				{-4, -2, -3, -5, -6, -3, -2, -4},
			},
			expectedError: nil,
		},
		{
			name: "chessboard.board is not set",
			board: [8][8]int{
				{7, 7, 7, 7, 7, 7, 7, 7},
				{7, 7, 7, 7, 7, 7, 7, 7},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{-7, -7, -7, -7, -7, -7, -7, -7},
				{-7, -7, -7, -7, -7, -7, -7, -7},
			},
			expectedResponse: [8][8]int{},
			expectedError: errors.New("chessboard.board is not set"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			entity := NewChessboardEntity(tc.board, "fen", "", "", "", 0, 0)
			// Act
			response, err := entity.GetBoard()
			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}